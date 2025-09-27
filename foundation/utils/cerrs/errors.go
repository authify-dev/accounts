package cerrs

type CustomErrorInterface interface {
	Error() string
	ToMap() map[string]interface{}
	GetCode() int
}

type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Scope   string `json:"scope"`
}

func NewCustomError(code int, message string, scope string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
		Scope:   scope,
	}
}

func (e *CustomError) Error() string {
	return e.Message
}

func (e *CustomError) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
		"scope":   e.Scope,
	}
}

func (e *CustomError) GetCode() int {
	return e.Code
}
