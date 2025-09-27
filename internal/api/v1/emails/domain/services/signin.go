package services

import (
	"accounts/internal/api/v1/emails/domain/entities"
	"context"
	"foundation/domain/criteria"
	"foundation/utils"
	"foundation/utils/cerrs"
	"net/http"
)

func (s *EmailsService) SignIn(
	ctx context.Context,
	entity entities.SignIn,
) utils.Response[entities.SignInResponse] {

	// Verificar el Email
	// Obtenemos el login
	// Generamos Tokens

	criteria_email := criteria.Criteria{
		Filters: *criteria.NewFilters(
			[]criteria.Filter{
				{
					Field:    "email",
					Value:    entity.Email,
					Operator: criteria.OperatorEqual,
				},
			},
		),
	}

	emails, err := s.repository.Matching(criteria_email)
	if err != nil {
		return utils.Response[entities.SignInResponse]{
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error getting email", "emails.signin.error_getting_email"),
			StatusCode: 500,
		}
	}

	if len(emails) == 0 {
		return utils.Response[entities.SignInResponse]{
			Error:      cerrs.NewCustomError(http.StatusNotFound, "email not found", "emails.signin.error_email_not_found"),
			StatusCode: 404,
		}
	}
	email := emails[0]

	// ----------------Verificar el Email----------------

	ok := s.password_controller.CheckPassword(entity.Password, email.Password)
	if !ok {
		return utils.Response[entities.SignInResponse]{
			Error:      cerrs.NewCustomError(http.StatusUnauthorized, "invalid password", "emails.signin.error_invalid_password"),
			StatusCode: 401,
		}
	}

	// ----------------Obtener el login----------------

	criteria_login := criteria.Criteria{
		Filters: *criteria.NewFilters(
			[]criteria.Filter{
				{
					Field:    "user_id",
					Value:    email.UserID,
					Operator: criteria.OperatorEqual,
				},
				{
					Field:    "entity_id",
					Value:    email.ID,
					Operator: criteria.OperatorEqual,
				},
				{
					Field:    "entity_type",
					Value:    "email",
					Operator: criteria.OperatorEqual,
				},
			},
		),
	}

	logins, err := s.login_methods_repository.Matching(criteria_login)
	if err != nil {
		return utils.Response[entities.SignInResponse]{
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error getting login", "emails.signin.error_getting_login"),
			StatusCode: 500,
		}
	}

	if len(logins) == 0 {
		return utils.Response[entities.SignInResponse]{
			Error:      cerrs.NewCustomError(http.StatusNotFound, "login not found", "emails.signin.error_login_not_found"),
			StatusCode: 404,
		}
	}
	login := logins[0]

	if !login.IsVerify {
		return utils.Response[entities.SignInResponse]{
			Error:      cerrs.NewCustomError(http.StatusUnauthorized, "email not verified", "emails.signin.error_email_not_verified"),
			StatusCode: 401,
		}
	}

	// ----------------Generar Tokens----------------
	refreshs_result := s.createRefreshToken(ctx, login)

	if refreshs_result.Err != nil {
		return utils.Response[entities.SignInResponse]{
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error creating refresh token"+refreshs_result.Err.Error(), "emails.signin.error_creating_refresh_token"),
			StatusCode: 500,
		}
	}

	result := s.generateTokens(ctx, login, refreshs_result.Data)

	if result.Err != nil {
		return utils.Response[entities.SignInResponse]{
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error generating tokens"+result.Err.Error(), "emails.signin.error_generating_tokens"),
			StatusCode: 500,
		}
	}

	s.publishActivationUserEvent(entity.Email, entity.Email)

	return utils.Response[entities.SignInResponse]{
		StatusCode: 200,
		Data: entities.SignInResponse{
			JWT:          result.Data.jwt,
			RefreshToken: result.Data.refresh_token,
		},
	}

}
