package services

import "accounts/internal/api/v1/oauth_logins/domain/repositories"

type OAuthService struct {
	google_repository repositories.OauthClientRepository
	oauth_repository  repositories.OAuthLoginRepository
}

func NewOAuthService(
	google_repository repositories.OauthClientRepository,
	oauth_repository repositories.OAuthLoginRepository,
) *OAuthService {
	return &OAuthService{
		google_repository: google_repository,
		oauth_repository:  oauth_repository,
	}
}
