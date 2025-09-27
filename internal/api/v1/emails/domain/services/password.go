package services

import (
	"accounts/internal/api/v1/emails/domain/commands"
	"accounts/internal/api/v1/emails/domain/entities"
	email_events "accounts/internal/api/v1/emails/domain/events"
	"accounts/internal/common/logger"
	"accounts/internal/core/domain"
	"accounts/internal/core/domain/event"
	"foundation/domain/criteria"
	"foundation/domain/customctx"
	"foundation/utils"
	"foundation/utils/cerrs"
	"log"
	"net/http"

	codes_entities "accounts/internal/api/v1/codes/domain/entities"
)

func (s *EmailsService) ResetPassword(
	cc *customctx.CustomContext,
	command commands.ResetPassword,
) utils.Response[entities.ResendActivationCodeResponse] {

	// Logger
	entry := logger.FromContext(cc.Context())
	entry.Info("Resend Code activation")

	// Get email entity
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
		return utils.Response[entities.ResendActivationCodeResponse]{
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error getting email", "emails.reset_password.error_getting_email"),
			StatusCode: 500,
		}
	}

	if len(emails) == 0 {
		return utils.Response[entities.ResendActivationCodeResponse]{
			Error:      cerrs.NewCustomError(http.StatusNotFound, "email not found", "emails.reset_password.error_email_not_found"),
			StatusCode: 404,
		}
	}

	email := emails[0]
	// Get user by email

	user, err := s.user_repository.Search(email.UserID)
	if err != nil {
		return utils.Response[entities.ResendActivationCodeResponse]{
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error getting user", "emails.reset_password.error_getting_user"),
			StatusCode: 500,
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
				{
					Field:    "organization_id",
					Value:    command.OrganizationID,
					Operator: criteria.OperatorEqual,
				},
			},
		),
	}

	codes, err := s.codes_repository.Matching(criteria_codes)
	if err != nil {
		return utils.Response[entities.ResendActivationCodeResponse]{
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error getting codes", "emails.reset_password.error_getting_codes"),
			StatusCode: 500,
		}
	}

	for _, code := range codes {
		if code.UserID == email.UserID && !code.IsRemoved {
			// Update code
			s.codes_repository.UpdateByFields(code.ID.String(), map[string]interface{}{
				"is_removed": true,
			})
		}
	}

	// Create new Code

	code := codes_entities.Code{
		UserID:         email.UserID,
		Entity:         domain.Entity{},
		Code:           generateCode(6),
		Type:           "reset_password",
		OrganizationID: command.OrganizationID,
	}

	result := s.codes_repository.Save(code)
	if result.Err != nil {
		return utils.Response[entities.ResendActivationCodeResponse]{
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error saving code", "emails.reset_password.error_saving_code"),
			StatusCode: 500,
		}
	}

	code.ID = result.Data.ID

	// Publish event

	s.publishResetPasswordEvent(command.Email, user.UserName, code.Code)

	entry.Info("Event published")

	return utils.Response[entities.ResendActivationCodeResponse]{
		StatusCode: 200,
		Data: entities.ResendActivationCodeResponse{
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
