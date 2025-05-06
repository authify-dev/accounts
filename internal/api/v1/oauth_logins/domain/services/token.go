package services

import (
	"accounts/internal/common/logger"
	"accounts/internal/utils"
	"context"

	"github.com/gin-gonic/gin"
)

func (s *OAuthService) GetTokenGoogle(ctx context.Context, code string) utils.Responses[map[string]any] {
	entry := logger.FromContext(ctx)
	entry.Info("GetTokenGoogle")

	result := s.google_repository.GetToken(code)
	if result.Err != nil {
		return utils.Responses[map[string]any]{
			Err:        result.Err,
			StatusCode: 500,
			Success:    false,
		}
	}

	return utils.Responses[map[string]any]{
		Body:       gin.H{"access_token": result.Data},
		StatusCode: 200,
		Success:    true,
	}
}
