package services

import (
	"accounts/internal/api/v1/emails/domain/entities"
	"accounts/internal/common/logger"
	"accounts/internal/core/domain"
	"accounts/internal/core/domain/criteria"
	"accounts/internal/utils"
	"context"
	"math/rand"
	"time"

	codes_entities "accounts/internal/api/v1/codes/domain/entities"
)

func (s *EmailsService) ResendActivationCode(
	ctx context.Context,
	entity entities.ResendActivationCode,
) utils.Responses[entities.ResendActivationCodeResponse] {

	// Logger
	entry := logger.FromContext(ctx)
	entry.Info("Resend Code activation")

	// Get email entity
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
		return utils.Responses[entities.ResendActivationCodeResponse]{
			StatusCode: 500,
			Errors:     []string{err.Error()},
		}
	}

	if len(emails) == 0 {
		return utils.Responses[entities.ResendActivationCodeResponse]{
			StatusCode: 404,
			Errors:     []string{"email not found"},
		}
	}

	email := emails[0]
	// Get user by email

	user, err := s.user_repository.Search(email.UserID)
	if err != nil {
		return utils.Responses[entities.ResendActivationCodeResponse]{
			StatusCode: 500,
			Errors:     []string{err.Error()},
		}
	}

	// Get Code
	criteria_codes := criteria.Criteria{
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
			},
		),
	}

	codes, err := s.codes_repository.Matching(criteria_codes)
	if err != nil {
		return utils.Responses[entities.ResendActivationCodeResponse]{
			StatusCode: 500,
			Errors:     []string{err.Error()},
		}
	}

	if len(codes) == 0 {
		return utils.Responses[entities.ResendActivationCodeResponse]{
			StatusCode: 404,
			Errors:     []string{"Not Code by Activation account is unused"},
		}

	}
	code := codes[0]

	// Update codes
	s.codes_repository.UpdateByFields(code.ID, map[string]interface{}{
		"is_removed": true,
		"user_id":    email.UserID,
	},
	)

	// Create new Code

	code = codes_entities.Code{
		UserID: email.UserID,
		Entity: domain.Entity{},
		Code:   generateCode(6),
	}

	result := s.codes_repository.Save(code)
	if result.Err != nil {
		return utils.Responses[entities.ResendActivationCodeResponse]{
			StatusCode: 500,
			Errors:     []string{result.Err.Error()},
		}
	}

	code.ID = result.Data

	// Publish event

	s.publishEvent(entity.Email, user.UserName, code.Code)

	entry.Info("Event published")

	return utils.Responses[entities.ResendActivationCodeResponse]{
		StatusCode: 200,
		Body: entities.ResendActivationCodeResponse{
			Message: "Activation code sent",
		},
	}
}

func generateCode(longitud int) string {
	const numeros = "0123456789"
	resultado := make([]byte, longitud)

	// Se crea un generador local de números aleatorios con semilla.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < longitud; i++ {
		resultado[i] = numeros[r.Intn(len(numeros))]
	}
	return string(resultado)
}
