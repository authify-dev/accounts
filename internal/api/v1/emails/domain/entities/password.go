package entities

type ResetPassword struct {
	Email string `json:"email" binding:"required,email"`
}
