package controllers

import "accounts/internal/api/v1/roles/domain/services"

// RolesController estructura para manejar la ruta de Health
type RolesController struct {
	roles_service services.RolesService
}

// NewRolesController constructor para RolesController
func NewRolesController(
	roles_service services.RolesService,
) *RolesController {
	return &RolesController{
		roles_service: roles_service,
	}
}
