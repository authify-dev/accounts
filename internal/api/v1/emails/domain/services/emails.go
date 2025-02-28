package services

import (
	codes "accounts/internal/api/v1/codes/domain/repositories"
	emails "accounts/internal/api/v1/emails/domain/repositories"
	login_methods "accounts/internal/api/v1/login_methods/domain/repositories"
	refresh "accounts/internal/api/v1/refresh_tokens/domain/repositories"
	roles "accounts/internal/api/v1/roles/domain/repositories"
	users "accounts/internal/api/v1/users/domain/repositories"
	"accounts/internal/common/controllers"
)

type EmailsService struct {
	repository               emails.EmailRepository
	user_repository          users.UserRepository
	role_repository          roles.RoleRepository
	login_methods_repository login_methods.LoginMethodRepository
	codes_repository         codes.CodeRepository
	refresh_repository       refresh.RefreshTokenRepository
	jwt_controller           controllers.JWTController
}

func NewEmailsService(
	repository emails.EmailRepository,
	user_repository users.UserRepository,
	role_repository roles.RoleRepository,
	login_methods_repository login_methods.LoginMethodRepository,
	codes_repository codes.CodeRepository,
	refresh_repository refresh.RefreshTokenRepository,
	jwt_contrller controllers.JWTController,
) *EmailsService {
	return &EmailsService{
		repository:               repository,
		user_repository:          user_repository,
		role_repository:          role_repository,
		login_methods_repository: login_methods_repository,
		codes_repository:         codes_repository,
		refresh_repository:       refresh_repository,
		jwt_controller:           jwt_contrller,
	}
}
