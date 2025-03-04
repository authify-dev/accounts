package services

import (
	"accounts/internal/api/v1/emails/domain/entities"
	"accounts/internal/core/domain/criteria"
	"accounts/internal/utils"
	"context"
	"fmt"
	"time"
)

func (s *EmailsService) Activate(
	ctx context.Context,
	entity entities.Activate,
) utils.Responses[entities.ActivateResponse] {

	// TODO: añadir el filed used_at al code y tipo de code

	// buscar el id del email
	// buscar el id del login con el email
	// buscar el code del login
	// comparar el code del login con el code del request

	// Generar JWT
	// Generar RefreshToken

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
		return utils.Responses[entities.ActivateResponse]{
			StatusCode: 500,
			Errors:     []string{err.Error()},
		}
	}

	if len(emails) == 0 {
		return utils.Responses[entities.ActivateResponse]{
			StatusCode: 404,
			Errors:     []string{"email not found"},
		}
	}
	email := emails[0]

	criteria_login := criteria.Criteria{
		Filters: *criteria.NewFilters(
			[]criteria.Filter{
				{
					Field:    "entity_id",
					Value:    email.ID,
					Operator: criteria.OperatorEqual,
				},
			},
		),
	}

	logins, err := s.login_methods_repository.Matching(criteria_login)
	if err != nil {
		return utils.Responses[entities.ActivateResponse]{
			StatusCode: 500,
			Errors:     []string{err.Error()},
		}
	}

	if len(logins) == 0 {
		return utils.Responses[entities.ActivateResponse]{
			StatusCode: 404,
			Errors:     []string{"login not found"},
		}
	}

	login := logins[0]
	fmt.Println(login)

	criteria_code := criteria.Criteria{
		Filters: *criteria.NewFilters(
			[]criteria.Filter{
				{
					Field:    "user_id",
					Value:    email.UserID,
					Operator: criteria.OperatorEqual,
				},
				{
					Field:    "is_removed",
					Value:    false,
					Operator: criteria.OperatorEqual,
				},
				{
					Field:    "created_at",
					Value:    time.Now().Add(-time.Minute * 15),
					Operator: criteria.OperatorGreaterThan,
				},
			},
		),
	}

	codes, err := s.codes_repository.Matching(criteria_code)
	if err != nil {
		return utils.Responses[entities.ActivateResponse]{
			StatusCode: 500,
			Errors:     []string{err.Error()},
		}
	}

	if len(codes) == 0 {
		return utils.Responses[entities.ActivateResponse]{
			StatusCode: 404,
			Errors:     []string{"code not found"},
		}
	}

	code := codes[0]

	if code.Code != entity.Code {
		return utils.Responses[entities.ActivateResponse]{
			StatusCode: 400,
			Errors:     []string{"invalid code"},
		}
	}

	return utils.Responses[entities.ActivateResponse]{
		StatusCode: 200,
		Body: entities.ActivateResponse{
			JWT:          "jwt",
			RefreshToken: "refresh",
		},
	}
}
