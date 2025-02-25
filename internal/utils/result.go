package utils

// --------------------------------
// UTILS
// --------------------------------
// Utils
//--------------------------------

type Result[R any] struct {
	Data R
	Err  error
}

type Responses[R any] struct {
	Body       R     `json:"body,omitempty"`
	Err        error `json:"error,omitempty"`
	StatusCode int   `json:"status_code" default:"200"`
}
