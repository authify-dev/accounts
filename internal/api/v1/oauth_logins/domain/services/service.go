package services

import (
	login_methods "accounts/internal/api/v1/login_methods/domain/repositories"
	oauth_logins "accounts/internal/api/v1/oauth_logins/domain/repositories"
	refresh_tokens "accounts/internal/api/v1/refresh_tokens/domain/repositories"
	roles "accounts/internal/api/v1/roles/domain/repositories"
	users "accounts/internal/api/v1/users/domain/repositories"
	"accounts/internal/common/controllers"
	"accounts/internal/core/domain/event"
)

type OAuthService struct {
	google_repository       oauth_logins.OauthClientRepository
	oauth_repository        oauth_logins.OAuthLoginRepository
	user_repository         users.UserRepository
	role_repository         roles.RoleRepository
	login_method_repository login_methods.LoginMethodRepository
	refresh_repository      refresh_tokens.RefreshTokenRepository
	jwt_controller          controllers.JWTController
	event_bus               event.EventBus
}

func NewOAuthService(
	google_repository oauth_logins.OauthClientRepository,
	oauth_repository oauth_logins.OAuthLoginRepository,
	user_repository users.UserRepository,
	role_repository roles.RoleRepository,
	login_method_repository login_methods.LoginMethodRepository,
	refresh_repository refresh_tokens.RefreshTokenRepository,
	jwt_controller controllers.JWTController,
	event_bus event.EventBus,
) *OAuthService {
	return &OAuthService{
		google_repository:       google_repository,
		oauth_repository:        oauth_repository,
		user_repository:         user_repository,
		role_repository:         role_repository,
		login_method_repository: login_method_repository,
		refresh_repository:      refresh_repository,
		jwt_controller:          jwt_controller,
		event_bus:               event_bus,
	}
}
