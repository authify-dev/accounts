package services

import (
	login_methods_repository "accounts/internal/api/v1/login_methods/domain/repositories"
	"accounts/internal/api/v1/refresh_tokens/domain/repositories"

	jwt_controller "accounts/internal/common/controllers"
)

type RefreshTokensService struct {
	repository               repositories.RefreshTokenRepository
	login_methods_repository login_methods_repository.LoginMethodRepository
	jwt_controller           jwt_controller.JWTController
}

func NewRefreshTokensService(
	repository repositories.RefreshTokenRepository,
	login_methods_repository login_methods_repository.LoginMethodRepository,
	jwt_controller jwt_controller.JWTController,
) *RefreshTokensService {
	return &RefreshTokensService{
		repository:               repository,
		login_methods_repository: login_methods_repository,
		jwt_controller:           jwt_controller,
	}
}
