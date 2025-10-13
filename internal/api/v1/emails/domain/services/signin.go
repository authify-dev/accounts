package services

import (
	"accounts/internal/api/v1/emails/domain/commands"
	"accounts/internal/api/v1/emails/domain/entities"
	"accounts/internal/common/logger"
	"foundation/domain/criteria"
	"foundation/domain/customctx"
	"foundation/utils"
	"foundation/utils/cerrs"
	"net/http"
)

func (s *EmailsService) SignIn(
	cc *customctx.CustomContext,
	command commands.SignIn,
) utils.Response[entities.SignInResponse] {

	entry := logger.FromContext(cc.Context())

	entry.Info("SignIn")

	criteria_email := criteria.Criteria{
		Filters: *criteria.NewFilters(
			[]criteria.Filter{
				{
					Field:    "email",
					Value:    command.Email,
					Operator: criteria.OperatorEqual,
				},
				{
					Field:    "organization_id",
					Value:    command.OrganizationID,
					Operator: criteria.OperatorEqual,
				},
			},
		),
	}

	emails, err := s.repository.Matching(criteria_email)
	if err != nil {
		entry.Error("Error getting email", err)
		return utils.Response[entities.SignInResponse]{
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error getting email", "emails.signin.error_getting_email"),
			StatusCode: 500,
		}
	}

	if len(emails) == 0 {
		entry.Error("Email not found")
		return utils.Response[entities.SignInResponse]{
			Error:      cerrs.NewCustomError(http.StatusNotFound, "email not found", "emails.signin.error_email_not_found"),
			StatusCode: 404,
		}
	}
	email := emails[0]

	// ----------------Verificar el Email----------------

	ok := s.password_controller.CheckPassword(command.Password, email.Password)
	if !ok {
		entry.Error("Invalid password")
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
				{
					Field:    "organization_id",
					Value:    command.OrganizationID,
					Operator: criteria.OperatorEqual,
				},
			},
		),
	}

	logins, err := s.login_methods_repository.Matching(criteria_login)
	if err != nil {
		entry.Error("Error getting login", err)
		return utils.Response[entities.SignInResponse]{
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error getting login", "emails.signin.error_getting_login"),
			StatusCode: 500,
		}
	}

	if len(logins) == 0 {
		entry.Error("Login not found")
		return utils.Response[entities.SignInResponse]{
			Error:      cerrs.NewCustomError(http.StatusNotFound, "login not found", "emails.signin.error_login_not_found"),
			StatusCode: 404,
		}
	}
	login := logins[0]

	if !login.IsVerify {
		entry.Error("Email not verified")
		return utils.Response[entities.SignInResponse]{
			Error:      cerrs.NewCustomError(http.StatusUnauthorized, "email not verified", "emails.signin.error_email_not_verified"),
			StatusCode: 401,
		}
	}

	// ----------------Generar Tokens----------------
	refreshs_result := s.createRefreshToken(cc.Context(), login)

	if refreshs_result.Err != nil {
		entry.Error("Error creating refresh token", refreshs_result.Err)
		return utils.Response[entities.SignInResponse]{
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error creating refresh token"+refreshs_result.Err.Error(), "emails.signin.error_creating_refresh_token"),
			StatusCode: 500,
		}
	}

	result := s.generateTokens(cc.Context(), login, refreshs_result.Data)

	if result.Err != nil {
		entry.Error("Error generating tokens", result.Err)
		return utils.Response[entities.SignInResponse]{
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error generating tokens"+result.Err.Error(), "emails.signin.error_generating_tokens"),
			StatusCode: 500,
		}
	}

	s.publishActivationUserEvent(command.Email, command.Email)

	return utils.Response[entities.SignInResponse]{
		StatusCode: 200,
		Data: entities.SignInResponse{
			JWT:          result.Data.jwt,
			RefreshToken: result.Data.refresh_token,
		},
	}

}
