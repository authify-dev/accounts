package services

import (
	emails "accounts/internal/api/v1/emails/domain/repositories"
	login_methods "accounts/internal/api/v1/login_methods/domain/repositories"
	roles "accounts/internal/api/v1/roles/domain/repositories"
	users "accounts/internal/api/v1/users/domain/repositories"
)

type EmailsService struct {
	repository               emails.EmailRepository
	user_repository          users.UserRepository
	role_repository          roles.RoleRepository
	login_methods_repository login_methods.LoginMethodRepository
}

func NewEmailsService(
	repository emails.EmailRepository,
	user_repository users.UserRepository,
	role_repository roles.RoleRepository,
	login_methods_repository login_methods.LoginMethodRepository,
) *EmailsService {
	return &EmailsService{
		repository:               repository,
		user_repository:          user_repository,
		role_repository:          role_repository,
		login_methods_repository: login_methods_repository,
	}
}
