package services

import (
	"accounts/internal/common/logger"
	"accounts/internal/utils"
	"context"
)

func (s *OAuthService) GetLinkGoogle(ctx context.Context) utils.Responses[string] {

	entry := logger.FromContext(ctx)
	entry.Info("GetLinkGoogle")

	result := s.google_repository.GetLink()
	if result.Err != nil {
		return utils.Responses[string]{
			Err:        result.Err,
			StatusCode: 500,
			Success:    false,
		}
	}

	return utils.Responses[string]{
		Body:       result.Data,
		StatusCode: 200,
		Success:    true,
	}
}
