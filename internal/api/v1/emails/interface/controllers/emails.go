package controllers

import "accounts/internal/api/v1/emails/domain/services"

// EmailsController estructura para manejar la ruta de Health
type EmailsController struct {
	userService services.EmailsService
}

// NewEmailsController constructor para EmailsController
func NewEmailsController(
	userService services.EmailsService,
) *EmailsController {
	return &EmailsController{
		userService: userService,
	}
}
