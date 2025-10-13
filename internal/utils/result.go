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

type Either[R any] struct {
	Data R
	Err  error
}

type Responses[R any] struct {
	Body       R        `json:"body,omitempty"`
	Err        error    `json:"error,omitempty"`
	StatusCode int      `json:"status_code" default:"200"`
	Errors     []string `json:"errors,omitempty"`
	Success    bool     `json:"success" default:"true"`
}

func (r Responses[R]) ToMap() map[string]interface{} {
	responseMap := make(map[string]interface{})
	r.Success = true

	if len(r.Errors) > 0 {
		responseMap["errors"] = r.Errors
		if r.Err != nil {
			responseMap["error"] = r.Err.Error()
		}
		r.Success = false
	} else {
		responseMap["body"] = r.Body
		if r.Err != nil {
			responseMap["error"] = r.Err.Error()
			r.Success = false
		}
	}

	responseMap["success"] = r.Success

	return responseMap
}
