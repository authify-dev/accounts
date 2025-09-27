package logger

import "fmt"

// LogFields represents the structure for logging fields.
// It includes information such as the request ID, method, client IP, and user ID.
type LogFields struct {
	TraceID  string `json:"trace_id"`
	Method   string `json:"method"`
	ClientIP string `json:"client_ip"`
	UserID   string `json:"user_id"`
	CallerID string `json:"caller_id"`
	Path     string `json:"path"`
	CallIn   string `json:"call_in"`
}

func (f LogFields) ToString() string {
	return fmt.Sprintf(" trace_id:%s | caller_id:%s | method:%s | client_ip:%s | user_id:%s ",
		f.TraceID,
		f.CallerID,
		f.Method,
		f.ClientIP,
		f.UserID,
	)
}
