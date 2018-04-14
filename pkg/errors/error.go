package errors

type ErrorReason struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// APIError is the error type usually returned by functions in the server
// It describes the code, message, and recoded time when an error occurred.
type APIError struct {
	ErrorReason ErrorReason `json:"error"`
	StatusCode  int         `json:"status_code"`
	RecordedAt  string      `json:"recorded_at"`
	Err         error
}

func (e *APIError) Error() string {
	if e == nil {
		return "<nil>"
	}

	er := e.ErrorReason
	s := ""
	if er.Code != "" {
		s = "API " + er.Code
	}

	if er.Message != "" {
		s = s + ":" + er.Message
	}

	return s
}
