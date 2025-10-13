package services

import (
	codes "accounts/internal/api/v1/codes/domain/entities"
	"accounts/internal/api/v1/emails/domain/entities"
	email_events "accounts/internal/api/v1/emails/domain/events"
	"accounts/internal/api/v1/emails/domain/steps"
	"fmt"
	"net/http"
	"strings"

	"accounts/internal/core/domain/event"
	"log"

	users "accounts/internal/api/v1/users/domain/entities"

	"accounts/internal/common/controllers/saga"
	"accounts/internal/common/logger"
	"context"
	"foundation/utils"
	"foundation/utils/cerrs"
)

func (s EmailsService) generateUser(entity entities.SignUp) utils.Result[users.User] {
	return utils.Result[users.User]{Data: users.User{
		UserName:       entity.UserName,
		Role:           entity.Role,
		OrganizationID: entity.OrganizationID,
	}}
}

func (s EmailsService) generateEmail(entity entities.SignUp) utils.Result[entities.Email] {

	password_hashed, err := s.password_controller.HashPassword(entity.Password)
	if err != nil {
		return utils.Result[entities.Email]{Err: cerrs.NewCustomError(http.StatusInternalServerError, "error hashing password", "emails.signup.error_hashing_password")}
	}

	return utils.Result[entities.Email]{Data: entities.Email{
		Email:          entity.Email,
		Password:       password_hashed,
		OrganizationID: entity.OrganizationID,
	}}
}

type RegisterUserFlow struct {
	controller *saga.SAGA_Controller
	results    map[string]utils.Result[any]
}

func (s EmailsService) registerUserWithEmail(ctx context.Context, entity entities.SignUp) utils.Result[RegisterUserFlow] {
	// Generar usuario
	user_result := s.generateUser(entity)
	if user_result.Err != nil {
		return utils.Result[RegisterUserFlow]{Err: user_result.Err}
	}
	user := user_result.Data

	// Generar email
	email_result := s.generateEmail(entity)
	if email_result.Err != nil {
		return utils.Result[RegisterUserFlow]{Err: user_result.Err}
	}
	email := email_result.Data

	// Crear usuario y email
	controller := saga.SAGA_Controller{
		Steps: []saga.SAGA_Step[any]{
			steps.NewCreateUserStep(
				s.user_repository,
				s.role_repository,
				user,
			),
			steps.NewCreateEmailStep(
				s.user_repository,
				s.repository,
				email,
			),
		},
	}

	results := controller.Executed(ctx)

	if !controller.Ok() {
		return utils.Result[RegisterUserFlow]{Err: cerrs.NewCustomError(http.StatusInternalServerError, "error creating user and email"+strings.Join(controller.Errors(), ","), "emails.signup.error_creating_user_and_email")}
	}

	return utils.Result[RegisterUserFlow]{Data: RegisterUserFlow{
		controller: &controller,
		results:    results,
	}}
}

type RegisterLoginFlow struct {
	controller *saga.SAGA_Controller
	results    map[string]utils.Result[any]
}

func (s EmailsService) registerLogin(ctx context.Context, registerUserFlow RegisterUserFlow) utils.Result[RegisterLoginFlow] {
	results := registerUserFlow.results
	controller := registerUserFlow.controller

	email := results["entities.Email"].Data.(entities.Email)

	user := results["entities.User"].Data.(users.User)

	// Crear login, refresh token y code
	controller_login := saga.SAGA_Controller{
		Steps: []saga.SAGA_Step[any]{
			steps.NewCreateLoginStep(
				s.login_methods_repository,
				user.ID.String(),
				email.ID.String(),
				"email",
				user.OrganizationID,
			),
			steps.NewCreateCodeStep(
				s.codes_repository,
				user.ID.String(),
				user.OrganizationID,
			),
		},
		PrevSaga: controller,
	}

	results_login := controller_login.Executed(ctx)

	if !controller_login.Ok() {
		return utils.Result[RegisterLoginFlow]{Err: cerrs.NewCustomError(http.StatusInternalServerError, "error creating login, refresh token and code", "emails.signup.error_creating_login_refresh_token_and_code")}
	}

	return utils.Result[RegisterLoginFlow]{Data: RegisterLoginFlow{
		controller: &controller_login,
		results:    results_login,
	}}
}

type GenerateTokensFlow struct {
	jwt           string
	refresh_token string
}

func (s EmailsService) publishRegisteredUserEvent(email string, user_name string, code string) {

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

func (s *EmailsService) SignUp(
	ctx context.Context,
	entity entities.SignUp,
) utils.Response[entities.SignUpResponse] {

	// Logger
	entry := logger.FromContext(ctx)
	entry.Info("SignUp User by Email")

	// Crear usuario y email
	results_map := s.registerUserWithEmail(ctx, entity)

	if results_map.Err != nil {
		entry.Error(fmt.Sprintf("Error al crear usuario y email: %s", results_map.Err.Error()))
		return utils.Response[entities.SignUpResponse]{Error: results_map.Err, StatusCode: 500}
	}

	results := results_map.Data

	entry.Info("User and Email created")

	// Crear login, refresh token y code
	result_login := s.registerLogin(ctx, results)

	if result_login.Err != nil {
		entry.Error(fmt.Sprintf("Error al crear login, refresh token y code: %s", result_login.Err.Error()))
		return utils.Response[entities.SignUpResponse]{Error: result_login.Err, StatusCode: 500}
	}

	results_login := result_login.Data.results

	entry.Info("Login, Refresh Token and Code created")

	// Publicar evento
	code := results_login["entities.Code"].Data.(codes.Code)

	s.publishRegisteredUserEvent(entity.Email, entity.UserName, code.Code)

	entry.Info("Event published")

	// Response

	response := utils.Response[entities.SignUpResponse]{
		Data: entities.SignUpResponse{
			Message: "User created check your email by activate your account",
		},
	}

	return response
}
