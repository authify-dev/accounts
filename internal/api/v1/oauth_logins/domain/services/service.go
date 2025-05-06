package services

import "accounts/internal/api/v1/oauth_logins/domain/repositories"

type OAuthService struct {
	google_repository repositories.OauthClientRepository
}

func NewOAuthService(
	google_repository repositories.OauthClientRepository,
) *OAuthService {
	return &OAuthService{
		google_repository: google_repository,
	}
}
