package services

import (
	code_rep "accounts/internal/api/v1/codes/domain/repositories"
	pending_registrations_rep "accounts/internal/api/v1/pending_registrations/domain/repositories"
	roles_rep "accounts/internal/api/v1/roles/domain/repositories"
	users_rep "accounts/internal/api/v1/users/domain/repositories"
	"accounts/internal/core/domain/event"
)

type PendingRegistrationsService struct {
	repository      pending_registrations_rep.PendingRegistrationsRepository
	rolesRepository roles_rep.RoleRepository
	usersRepository users_rep.UserRepository
	codeRepository  code_rep.CodeRepository
	event_bus       event.EventBus
}

func NewPendingRegistrationsService(
	repository pending_registrations_rep.PendingRegistrationsRepository,
	rolesRepository roles_rep.RoleRepository,
	usersRepository users_rep.UserRepository,
	codeRepository code_rep.CodeRepository,
	eventBus event.EventBus,
) *PendingRegistrationsService {
	return &PendingRegistrationsService{
		repository:      repository,
		rolesRepository: rolesRepository,
		usersRepository: usersRepository,
		codeRepository:  codeRepository,
		event_bus:       eventBus,
	}
}
