package email_events

import "github.com/google/uuid"

type UserActivated struct {
	Email    string `json:"email"`
	UserName string `json:"user"`
}

func (u UserActivated) ToPrimitive() map[string]interface{} {
	return map[string]interface{}{
		"email": u.Email,
		"user":  u.UserName,
	}
}

func (u UserActivated) EventName() string {
	return "user.activated"
}

func (u UserActivated) AggregateID() string {
	return ""
}

func (u UserActivated) EventID() string {
	return uuid.New().String()
}
