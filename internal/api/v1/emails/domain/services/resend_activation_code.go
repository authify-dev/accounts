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
	"math/rand"
	"net/http"
	"time"

	codes_entities "accounts/internal/api/v1/codes/domain/entities"
)

func (s *EmailsService) ResendActivationCode(
	cc *customctx.CustomContext,
	command commands.ResendActivationCodeCommand,
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
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error getting email", "emails.resend_activation_code.error_getting_email"),
			StatusCode: 500,
		}
	}

	if len(emails) == 0 {
		return utils.Response[entities.ResendActivationCodeResponse]{
			Error:      cerrs.NewCustomError(http.StatusNotFound, "email not found", "emails.resend_activation_code.error_email_not_found"),
			StatusCode: 404,
		}
	}

	email := emails[0]
	// Get user by email

	user, err := s.user_repository.Search(email.UserID)
	if err != nil {
		return utils.Response[entities.ResendActivationCodeResponse]{
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error getting user", "emails.resend_activation_code.error_getting_user"),
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
					Value:    "activation",
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
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error getting codes", "emails.resend_activation_code.error_getting_codes"),
			StatusCode: 500,
		}
	}

	if len(codes) == 0 {
		return utils.Response[entities.ResendActivationCodeResponse]{
			Error:      cerrs.NewCustomError(http.StatusNotFound, "not code by activation account is unused", "emails.resend_activation_code.error_not_code_by_activation_account_is_unused"),
			StatusCode: 404,
		}

	}
	code := codes[0]

	// Update codes
	s.codes_repository.UpdateByFields(code.ID.String(), map[string]interface{}{
		"is_removed": true,
		"user_id":    email.UserID,
	},
	)

	// Create new Code

	code = codes_entities.Code{
		UserID:         email.UserID,
		Entity:         domain.Entity{},
		Code:           generateCode(6),
		OrganizationID: command.OrganizationID,
	}

	result := s.codes_repository.Save(code)
	if result.Err != nil {
		return utils.Response[entities.ResendActivationCodeResponse]{
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error saving code", "emails.resend_activation_code.error_saving_code"),
			StatusCode: 500,
		}
	}

	code.ID = result.Data.ID

	// Publish event

	s.publishResendCodeEvent(command.Email, user.UserName, code.Code)

	entry.Info("Event published")

	return utils.Response[entities.ResendActivationCodeResponse]{
		StatusCode: 200,
		Data: entities.ResendActivationCodeResponse{
			Message: "Activation code sent",
		},
	}
}

func (s EmailsService) publishResendCodeEvent(email string, user_name string, code string) {

	user_event := email_events.UserRegistered{
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
