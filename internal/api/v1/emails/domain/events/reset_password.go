package email_events

import "github.com/google/uuid"

type ResetPassword struct {
	Email            string `json:"email"`
	CodeVerification string `json:"code"`
	UserName         string `json:"user"`
}

func (u ResetPassword) ToPrimitive() map[string]interface{} {
	return map[string]interface{}{
		"email": u.Email,
		"code":  u.CodeVerification,
		"user":  u.UserName,
	}
}

func (u ResetPassword) EventName() string {
	return "user.reset_password"
}

func (u ResetPassword) AggregateID() string {
	return ""
}

func (u ResetPassword) EventID() string {
	return uuid.New().String()
}

type ChangedPassword struct {
	Email    string `json:"email"`
	UserName string `json:"user"`
}

func (u ChangedPassword) ToPrimitive() map[string]interface{} {
	return map[string]interface{}{
		"email": u.Email,
		"user":  u.UserName,
	}
}

func (u ChangedPassword) EventName() string {
	return "user.changed_password"
}

func (u ChangedPassword) AggregateID() string {
	return ""
}

func (u ChangedPassword) EventID() string {
	return uuid.New().String()
}
