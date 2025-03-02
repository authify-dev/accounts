package email_events

import "github.com/google/uuid"

type UserRegistered struct {
	Email            string
	CodeVerification string
	UserName         string
}

func (u UserRegistered) ToPrimitive() map[string]interface{} {
	return map[string]interface{}{
		"email": u.Email,
		"code":  u.CodeVerification,
		"user":  u.UserName,
	}
}

func (u UserRegistered) EventName() string {
	return "user.registered"
}

func (u UserRegistered) AggregateID() string {
	return ""
}

func (u UserRegistered) EventID() string {
	return uuid.New().String()
}
