package controllers

import "accounts/internal/api/v1/role_policies/domain/services"

type RolePoliciesController struct {
	service *services.RolePoliciesService
}

func NewRolePoliciesController(service *services.RolePoliciesService) *RolePoliciesController {
	return &RolePoliciesController{
		service: service,
	}
}
