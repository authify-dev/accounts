package controllers

import "accounts/internal/api/v1/refresh_tokens/domain/services"

type RefreshTokensController struct {
	service *services.RefreshTokensService
}

func NewRefreshTokensController(
	service *services.RefreshTokensService,
) *RefreshTokensController {
	return &RefreshTokensController{
		service: service,
	}
}
