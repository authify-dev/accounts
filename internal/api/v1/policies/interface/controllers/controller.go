package controllers

import "accounts/internal/api/v1/policies/domain/services"

type PoliciesController struct {
	policies_service *services.PoliciesService
}

func NewPoliciesController(policies_service *services.PoliciesService) *PoliciesController {
	return &PoliciesController{policies_service: policies_service}
}
