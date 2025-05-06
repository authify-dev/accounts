package services

import (
	"accounts/internal/common/logger"
	"accounts/internal/utils"
	"context"

	"github.com/gin-gonic/gin"
)

func (s *OAuthService) GetUserInfoGoogle(ctx context.Context, token string) utils.Responses[map[string]any] {
	entry := logger.FromContext(ctx)
	entry.Info("GetUserInfoGoogle")

	result := s.google_repository.GetUserInfo(token)
	if result.Err != nil {
		return utils.Responses[map[string]any]{
			Err:        result.Err,
			StatusCode: 500,
			Success:    false,
		}
	}

	return utils.Responses[map[string]any]{
		Body:       gin.H{"user_info": result.Data},
		StatusCode: 200,
		Success:    true,
	}
}
