package entities

type ResetPassword struct {
	Email string `json:"email" binding:"required,email"`
}

type ConfirmPassword struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=100"`
	Code     string `json:"code" binding:"required,min=6,max=6"`
}

type ResetPasswordResponse struct {
	Message string `json:"message"`
}
