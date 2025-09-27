package services

import (
	"accounts/internal/api/v1/emails/domain/entities"
	"accounts/internal/common/logger"
	"context"
	"foundation/domain/criteria"
	"foundation/utils"
	"foundation/utils/cerrs"
	"net/http"
)

func (s *RefreshTokensService) Create(
	ctx context.Context,
	refresh_token string,
) utils.Response[entities.SignInResponse] {

	entry := logger.FromContext(ctx)

	entry.Info("Creating new JWT")

	claim, err := s.jwt_controller.ValidateToken(ctx, refresh_token)
	if err != nil {
		entry.Error("Failed to validate token", err)
		return utils.Response[entities.SignInResponse]{
			StatusCode: 401,
			Error:      cerrs.NewCustomError(http.StatusUnauthorized, "Failed to validate token", "refresh_tokens.create.failed_to_validate_token"),
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
		return utils.Response[entities.SignInResponse]{
			StatusCode: 500,
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "Failed to get refresh token", "refresh_tokens.create.failed_to_get_refresh_token"),
		}
	}

	if len(refresh_ents) == 0 {
		entry.Error("Refresh token not found")
		return utils.Response[entities.SignInResponse]{
			StatusCode: 404,
			Error:      cerrs.NewCustomError(http.StatusNotFound, "Refresh token not found", "refresh_tokens.create.refresh_token_not_found"),
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
		return utils.Response[entities.SignInResponse]{
			StatusCode: 500,
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "Failed to get login method", "refresh_tokens.create.failed_to_get_login_method"),
		}
	}

	if len(login_ents) == 0 {
		entry.Error("Login method not found")
		return utils.Response[entities.SignInResponse]{
			StatusCode: 404,
			Error:      cerrs.NewCustomError(http.StatusNotFound, "Login method not found", "refresh_tokens.create.login_method_not_found"),
		}
	}

	login_method_entity := login_ents[0]

	jwt := login_method_entity.ToJWT(ctx, s.jwt_controller)

	return utils.Response[entities.SignInResponse]{
		StatusCode: 201,
		Data: entities.SignInResponse{
			JWT:          jwt,
			RefreshToken: refresh_token,
		},
	}
}
