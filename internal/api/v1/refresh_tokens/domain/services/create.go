package services

import (
	"accounts/internal/api/v1/emails/domain/entities"
	"accounts/internal/common/logger"
	"accounts/internal/core/domain/criteria"
	"accounts/internal/utils"
	"context"
)

func (s *RefreshTokensService) Create(
	ctx context.Context,
	refresh_token string,
) utils.Responses[entities.SignInResponse] {

	entry := logger.FromContext(ctx)

	entry.Info("Creating new JWT")

	claim, err := s.jwt_controller.ValidateToken(refresh_token)
	if err != nil {
		entry.Error("Failed to validate token", err)
		return utils.Responses[entities.SignInResponse]{
			StatusCode: 401,
			Err:        err,
		}
	}

	refersh_token_id := claim["id"].(string)

	// -----------------Obtener el refresh token entity----------------

	cri := criteria.Criteria{
		Filters: *criteria.NewFilters(
			[]criteria.Filter{
				{
					Field:    "id",
					Value:    refersh_token_id,
					Operator: criteria.OperatorEqual,
				},
			},
		),
	}

	refresh_ents, err := s.repository.Matching(cri)

	if err != nil {
		entry.Error("Failed to get refresh token", err)
		return utils.Responses[entities.SignInResponse]{
			StatusCode: 500,
			Errors:     []string{err.Error()},
		}
	}

	if len(refresh_ents) == 0 {
		entry.Error("Refresh token not found")
		return utils.Responses[entities.SignInResponse]{
			StatusCode: 404,
			Errors:     []string{"refresh token not found"},
		}
	}

	refresh_token_entity := refresh_ents[0]

	// -----------------Obtener el login auth----------------
	cri_login := criteria.Criteria{
		Filters: *criteria.NewFilters(
			[]criteria.Filter{
				{
					Field:    "user_id",
					Value:    refresh_token_entity.UserID,
					Operator: criteria.OperatorEqual,
				},
				{
					Field:    "id",
					Value:    refresh_token_entity.LoginMethodID,
					Operator: criteria.OperatorEqual,
				},
			},
		),
	}

	login_ents, err := s.login_methods_repository.Matching(cri_login)
	if err != nil {
		entry.Error("Failed to get login method", err)
		return utils.Responses[entities.SignInResponse]{
			StatusCode: 500,
			Errors:     []string{err.Error()},
		}
	}

	if len(login_ents) == 0 {
		entry.Error("Login method not found")
		return utils.Responses[entities.SignInResponse]{
			StatusCode: 404,
			Errors:     []string{"login method not found"},
		}
	}

	login_method_entity := login_ents[0]

	jwt := login_method_entity.ToJWT(s.jwt_controller)

	return utils.Responses[entities.SignInResponse]{
		StatusCode: 201,
		Body: entities.SignInResponse{
			JWT:          jwt,
			RefreshToken: refresh_token,
		},
	}
}
