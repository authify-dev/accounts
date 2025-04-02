package services

import (
	"accounts/internal/api/v1/emails/domain/entities"
	"accounts/internal/core/domain/criteria"
	"accounts/internal/utils"
	"context"
)

func (s *EmailsService) SignIn(
	ctx context.Context,
	entity entities.SignIn,
) utils.Responses[entities.SignInResponse] {

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
		return utils.Responses[entities.SignInResponse]{
			StatusCode: 500,
			Errors:     []string{err.Error()},
		}
	}

	if len(emails) == 0 {
		return utils.Responses[entities.SignInResponse]{
			StatusCode: 404,
			Errors:     []string{"email not found"},
		}
	}
	email := emails[0]

	// ----------------Verificar el Email----------------

	ok := s.password_controller.CheckPassword(entity.Password, email.Password)
	if !ok {
		return utils.Responses[entities.SignInResponse]{
			StatusCode: 401,
			Errors:     []string{"invalid password"},
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
		return utils.Responses[entities.SignInResponse]{
			StatusCode: 500,
			Errors:     []string{err.Error()},
		}
	}

	if len(logins) == 0 {
		return utils.Responses[entities.SignInResponse]{
			StatusCode: 404,
			Errors:     []string{"login not found"},
		}
	}
	login := logins[0]

	if !login.IsVerify {
		return utils.Responses[entities.SignInResponse]{
			StatusCode: 401,
			Errors:     []string{"email not verified"},
		}
	}

	// ----------------Generar Tokens----------------
	refreshs_result := s.createRefreshToken(ctx, login)

	if refreshs_result.Err != nil {
		return utils.Responses[entities.SignInResponse]{
			StatusCode: 500,
			Errors:     []string{refreshs_result.Err.Error()},
		}
	}

	result := s.generateTokens(login, refreshs_result.Data)

	if result.Err != nil {
		return utils.Responses[entities.SignInResponse]{
			StatusCode: 500,
			Errors:     []string{result.Err.Error()},
		}
	}

	s.publishActivationUserEvent(entity.Email, entity.Email)

	return utils.Responses[entities.SignInResponse]{
		StatusCode: 200,
		Body: entities.SignInResponse{
			JWT:          result.Data.jwt,
			RefreshToken: result.Data.refresh_token,
		},
	}

}
