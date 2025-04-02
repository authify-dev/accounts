package responses

type Response struct {
	Status   int         `json:"status"`
	Data     interface{} `json:"data,omitempty"`
	Metadata interface{} `json:"metadata,omitempty"`
	Errors   interface{} `json:"errors,omitempty"`
}
