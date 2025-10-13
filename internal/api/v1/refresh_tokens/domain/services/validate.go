package services

import (
	"accounts/internal/common/logger"
	"accounts/internal/utils"
	"context"
)

func (s *RefreshTokensService) Validate(
	ctx context.Context,
	refresh_token string,
) utils.Responses[map[string]interface{}] {

	entry := logger.FromContext(ctx)

	entry.Info("Creating new JWT")

	claim, err := s.jwt_controller.ValidateToken(ctx, refresh_token)
	if err != nil {
		entry.Error("Failed to validate token", err)
		return utils.Responses[map[string]interface{}]{
			StatusCode: 401,
			Errors:     []string{err.Error()},
		}
	}

	entity_type, ok := claim["entity_type"].(string)
	if !ok {
		entry.Error("entity_type key is missing or not a string in the token")
		return utils.Responses[map[string]interface{}]{
			StatusCode: 401,
			Errors:     []string{"entity_type key is missing or not a string in the token"},
		}
	}
	if entity_type == "" {
		entry.Error("Token is not a JWT")
		return utils.Responses[map[string]interface{}]{
			StatusCode: 401,
			Errors:     []string{"Token is not a JWT"},
		}
	}

	entry.Info("Token validated")
	entry.Info(claim)

	return utils.Responses[map[string]interface{}]{
		StatusCode: 200,
		Body: map[string]interface{}{
			"is_valid":    true,
			"entity_type": entity_type,
		},
	}
}
