package controllers

import "accounts/internal/api/v1/users/domain/services"

// UsersController estructura para manejar la ruta de Health
type UsersController struct {
	userService services.UsersService
}

// NewUsersController constructor para UsersController
func NewUsersController(
	userService services.UsersService,
) *UsersController {
	return &UsersController{
		userService: userService,
	}
}
