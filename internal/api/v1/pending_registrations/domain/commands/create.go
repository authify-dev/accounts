package commands

type CreatePendingRegistrationCommand struct {
	Email    string `json:"email" binding:"required"`
	UserName string `json:"user_name,omitempty"`
	Role     string `json:"role" binding:"required"`
}
