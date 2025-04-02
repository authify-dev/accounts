package controllers

import "accounts/internal/api/v1/roles/domain/services"

// RolesController estructura para manejar la ruta de Health
type RolesController struct {
	userService services.RolesService
}

// NewRolesController constructor para RolesController
func NewRolesController(
	userService services.RolesService,
) *RolesController {
	return &RolesController{
		userService: userService,
	}
}
