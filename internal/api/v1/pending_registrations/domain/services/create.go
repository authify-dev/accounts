package services

import (
	codes_entities "accounts/internal/api/v1/codes/domain/entities"
	email_events "accounts/internal/api/v1/emails/domain/events"
	"accounts/internal/api/v1/pending_registrations/domain/commands"
	pending_registrations_entities "accounts/internal/api/v1/pending_registrations/domain/entities"
	users_entities "accounts/internal/api/v1/users/domain/entities"
	"accounts/internal/common/logger"
	"accounts/internal/core/domain"
	"accounts/internal/core/domain/event"
	"context"
	"foundation/domain/criteria"
	"foundation/utils"
	"foundation/utils/cerrs"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func (s *PendingRegistrationsService) Create(ctx context.Context, command commands.CreatePendingRegistrationCommand) utils.Response[string] {

	entry := logger.FromContext(ctx)

	entry.Info("Creating pending registration")

	// Validar la exixtencia del role
	cri := criteria.Criteria{
		Filters: *criteria.NewFilters(
			[]criteria.Filter{
				{
					Field:    "name",
					Operator: criteria.OperatorEqual,
					Value:    command.Role,
				},
			},
		),
	}

	roles, err := s.rolesRepository.Matching(cri)
	if err != nil {
		entry.Error("Error getting roles", "error", err)
		return utils.Response[string]{
			StatusCode: http.StatusInternalServerError,
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "Error getting roles", "pending_registrations.create.error_getting_roles"),
		}
	}

	if len(roles) == 0 {
		entry.Error("Role not found")
		return utils.Response[string]{
			StatusCode: http.StatusNotFound,
			Error:      cerrs.NewCustomError(http.StatusNotFound, "Role not found", "pending_registrations.create.role_not_found"),
		}
	}

	// Crear el user

	user := users_entities.User{
		UserName: command.UserName,
		Role:     command.Role,
		RoleID:   roles[0].ID.String(),
	}

	result := s.usersRepository.Save(user)
	if result.Err != nil {
		entry.Error("Error saving user", "error", result.Err)
		return utils.Response[string]{
			StatusCode: http.StatusInternalServerError,
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "Error saving user", "pending_registrations.create.error_saving_user"),
		}
	}

	// Crear el code
	code := codes_entities.Code{
		Entity: domain.Entity{},
		Code:   generateCode(6),
		Type:   "registration",
		UserID: result.Data.ID.String(),
	}

	resultCode := s.codeRepository.Save(code)
	if resultCode.Err != nil {
		entry.Error("Error saving code", "error", resultCode.Err)
		return utils.Response[string]{
			StatusCode: http.StatusInternalServerError,
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "Error saving code", "pending_registrations.create.error_saving_code"),
		}
	}

	code.ID = resultCode.Data.ID

	// Crear el pending registration
	pendingRegistration := pending_registrations_entities.PendingRegistration{
		Email:    command.Email,
		UserName: command.UserName,
		Role:     command.Role,
		CodeID:   code.ID.String(),
	}

	resultPendingRegistration := s.repository.Save(pendingRegistration)
	if resultPendingRegistration.Err != nil {
		entry.Error("Error saving pending registration", "error", resultPendingRegistration.Err)
		return utils.Response[string]{
			StatusCode: http.StatusInternalServerError,
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "Error saving pending registration", "pending_registrations.create.error_saving_pending_registration"),
		}
	}

	pendingRegistration.ID = resultPendingRegistration.Data.ID

	entry.Info("Pending registration created successfully", "pendingRegistration", pendingRegistration)

	// Enviar el email
	s.publishValidateRegistrationEvent(command.Email, command.UserName, code.Code)

	return utils.Response[string]{
		Data:       "Check your email to validate your registration",
		StatusCode: http.StatusOK,
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

func (s PendingRegistrationsService) publishValidateRegistrationEvent(email string, user_name string, code string) {

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
