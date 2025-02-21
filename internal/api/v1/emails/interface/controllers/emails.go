package controllers

import "accounts/internal/api/v1/users/domain/services"

// EmailsController estructura para manejar la ruta de Health
type EmailsController struct {
	userService services.UsersService
}

// NewEmailsController constructor para EmailsController
func NewEmailsController(
	userService services.UsersService,
) *EmailsController {
	return &EmailsController{
		userService: userService,
	}
}
