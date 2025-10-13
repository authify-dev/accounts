package controllers

import "accounts/internal/api/v1/oauth_logins/domain/services"

type OAuthController struct {
	service *services.OAuthService
}

func NewOAuthController(
	service *services.OAuthService,
) *OAuthController {
	return &OAuthController{
		service: service,
	}
}
