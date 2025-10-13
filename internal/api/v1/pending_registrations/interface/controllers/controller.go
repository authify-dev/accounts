package controllers

import "accounts/internal/api/v1/pending_registrations/domain/services"

type PendingRegistrationsController struct {
	service *services.PendingRegistrationsService
}

func NewPendingRegistrationsController(
	service *services.PendingRegistrationsService,
) *PendingRegistrationsController {
	return &PendingRegistrationsController{
		service: service,
	}
}
