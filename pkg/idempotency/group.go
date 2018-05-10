package idempotency

import "sync"

// call is an in-flight or completed Once call
type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

// Group represents a class of work and forms a namespace in which
// units of work can be executed with duplicate suppression.
type Group struct {
	mu sync.Mutex       // protects m
	m  map[string]*call // lazily initialized
}

// Once executes and returns the results of the given function, making
// sure that only one execution for a given key happens until the
// key is explicitly forgotten. If a duplicate comes in, the duplicate
// caller waits for the original to complete and receives the same results.
func (g *Group) Once(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}
	c := new(call)
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

	c.val, c.err = fn()
	if c.err != nil {
		g.mu.Lock()
		delete(g.m, key)
		g.mu.Unlock()
	}
	c.wg.Done()

	return c.val, c.err
}

// Forget forgets a key, allowing the next call for the key to execute
// the function.
func (g *Group) Forget(key string) {
	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()
}