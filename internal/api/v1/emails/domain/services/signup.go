package services

import (
	codes "accounts/internal/api/v1/codes/domain/entities"
	"accounts/internal/api/v1/emails/domain/entities"
	email_events "accounts/internal/api/v1/emails/domain/events"
	"accounts/internal/api/v1/emails/domain/steps"
	logins "accounts/internal/api/v1/login_methods/domain/entities"
	refreshs "accounts/internal/api/v1/refresh_tokens/domain/entities"

	"accounts/internal/core/domain/event"
	"log"

	users "accounts/internal/api/v1/users/domain/entities"

	"accounts/internal/common/controllers/saga"
	"accounts/internal/utils"
	"context"
)

func (s *EmailsService) SignUp(
	ctx context.Context,
	entity entities.SignUp,
) utils.Responses[entities.SignUpResponse] {

	user := users.User{
		UserName: entity.UserName,
		Role:     entity.Role,
	}

	password_hashed, err := s.password_controller.HashPassword(entity.Password)
	if err != nil {
		return utils.Responses[entities.SignUpResponse]{Errors: []string{err.Error()}, StatusCode: 500}
	}

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
				entities.Email{
					Email:    entity.Email,
					Password: password_hashed,
				},
			),
		},
	}

	results := controller.Executed(ctx)

	if !controller.Ok() {
		return utils.Responses[entities.SignUpResponse]{Errors: controller.Errors(), StatusCode: 500}
	}

	email := results["entities.Email"].Data.(entities.Email)

	user = results["entities.User"].Data.(users.User)

	// Crear login, refresh token y code
	controller_login := saga.SAGA_Controller{
		Steps: []saga.SAGA_Step[any]{
			steps.NewCreateLoginStep(
				s.login_methods_repository,
				user.ID,
				email.ID,
			),
			steps.NewCreateRefreshTokenStep(
				s.refresh_repository,
				user.ID,
			),
			steps.NewCreateCodeStep(
				s.codes_repository,
				user.ID,
			),
		},
		PrevSaga: &controller,
	}

	results_login := controller_login.Executed(ctx)

	if !controller_login.Ok() {
		return utils.Responses[entities.SignUpResponse]{Errors: controller_login.Errors(), StatusCode: 500}
	}

	login := results_login["entities.LoginMethod"].Data.(logins.LoginMethod)

	jwt := login.ToJWT(s.jwt_controller)

	refresh := results_login["entities.RefreshToken"].Data.(refreshs.RefreshToken)

	refresh_token := refresh.ToJWT(s.jwt_controller)

	res := entities.SignUpResponse{
		JWT:          jwt,
		RefreshToken: refresh_token,
	}

	response := utils.Responses[entities.SignUpResponse]{Body: res, StatusCode: 201}

	for _, result := range results {
		if result.Err != nil {
			response.Errors = append(response.Errors, result.Err.Error())
		}
	}

	code := results_login["entities.Code"].Data.(codes.Code)

	user_event := email_events.UserRegistered{
		Email:            email.Email,
		CodeVerification: code.Code,
		UserName:         user.UserName,
	}

	// Agregar el mensaje a la cola "new-users"
	if err := s.event_bus.Publish([]event.DomainEvent{
		user_event,
	}); err != nil {
		log.Println("Error al publicar el evento new-users")
	}

	return response
}
