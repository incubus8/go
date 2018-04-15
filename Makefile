deps:
	git submodule update --init --recursive

lint:
	@go fmt github.com/incubus8/go/...
