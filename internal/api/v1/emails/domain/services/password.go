package services

import (
	"accounts/internal/api/v1/emails/domain/entities"
	email_events "accounts/internal/api/v1/emails/domain/events"
	"accounts/internal/common/logger"
	"accounts/internal/core/domain"
	"accounts/internal/core/domain/criteria"
	"accounts/internal/core/domain/event"
	"accounts/internal/utils"
	"context"
	"log"

	codes_entities "accounts/internal/api/v1/codes/domain/entities"
)

func (s *EmailsService) ResetPassword(
	ctx context.Context,
	entity entities.ResetPassword,
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
				{
					Field:    "type",
					Value:    "reset_password",
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

	for _, code := range codes {
		if code.UserID == email.UserID && !code.IsRemoved {
			// Update code
			s.codes_repository.UpdateByFields(code.ID, map[string]interface{}{
				"is_removed": true,
			})
		}
	}

	// Create new Code

	code := codes_entities.Code{
		UserID: email.UserID,
		Entity: domain.Entity{},
		Code:   generateCode(6),
		Type:   "reset_password",
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

	s.publishResetPasswordEvent(entity.Email, user.UserName, code.Code)

	entry.Info("Event published")

	return utils.Responses[entities.ResendActivationCodeResponse]{
		StatusCode: 200,
		Body: entities.ResendActivationCodeResponse{
			Message: "Activation code sent",
		},
	}
}

func (s EmailsService) publishResetPasswordEvent(email string, user_name string, code string) {

	user_event := email_events.ResetPassword{
		Email:            email,
		CodeVerification: code,
		UserName:         user_name,
	}

	// Agregar el mensaje a la cola "new-users"
	if err := s.event_bus.Publish([]event.DomainEvent{
		user_event,
	}); err != nil {
		log.Println("Error al publicar el evento new-users")
		log.Println(err)
	}
}
