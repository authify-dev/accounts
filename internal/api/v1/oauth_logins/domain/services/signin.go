package services

import (
	"accounts/internal/api/v1/emails/domain/entities"
	"accounts/internal/api/v1/emails/domain/steps"
	"accounts/internal/common/controllers/saga"
	"accounts/internal/common/logger"
	"accounts/internal/core/domain"
	"accounts/internal/core/domain/criteria"
	"accounts/internal/core/domain/event"
	"accounts/internal/utils"
	"context"
	"errors"
	"fmt"
	"log"

	email_events "accounts/internal/api/v1/emails/domain/events"

	oauth_steps "accounts/internal/api/v1/oauth_logins/domain/steps"

	login_methods "accounts/internal/api/v1/login_methods/domain/entities"
	oauth_logins "accounts/internal/api/v1/oauth_logins/domain/entities"
	refreshs "accounts/internal/api/v1/refresh_tokens/domain/entities"
	users "accounts/internal/api/v1/users/domain/entities"

	"github.com/google/uuid"
)

type GenerateTokensFlow struct {
	jwt           string
	refresh_token string
}

func (s *OAuthService) SignInGoogle(ctx context.Context, code, role string) utils.Responses[entities.SignInResponse] {

	entry := logger.FromContext(ctx)

	// Obtener el token
	token_result := s.google_repository.GetToken(code)
	if token_result.Err != nil {
		return utils.Responses[entities.SignInResponse]{
			Err:        token_result.Err,
			StatusCode: 500,
			Success:    false,
		}
	}
	// Obtener la data del usuario
	user_info_result := s.google_repository.GetUserInfo(token_result.Data)
	if user_info_result.Err != nil {
		return utils.Responses[entities.SignInResponse]{
			Err:        user_info_result.Err,
			StatusCode: 500,
			Success:    false,
		}
	}

	// Verificamos que el email del user no este en uso

	cri := criteria.Criteria{
		Filters: *criteria.NewFilters(
			[]criteria.Filter{
				{
					Field:    "email",
					Operator: criteria.OperatorEqual,
					Value:    user_info_result.Data.Email,
				},
			},
		),
	}

	oauth, err := s.oauth_repository.Matching(cri)
	if err != nil {
		entry.Error("error matching oauth")
		return utils.Responses[entities.SignInResponse]{
			Err:        err,
			StatusCode: 500,
		}
	}

	controller := saga.SAGA_Controller{}
	var results map[string]utils.Result[any]
	var results_login map[string]utils.Result[any]

	var oauth_ent oauth_logins.OAuthLogin
	var user_ent users.User
	var login_ent login_methods.LoginMethod
	var refreshs_ent refreshs.RefreshToken

	if len(oauth) == 0 {
		entry.Info("Creating new user and oauth")

		// Generar usuario
		user := users.User{
			UserName: utils.GenerateRandomUserName(),
			Role:     role,
			Name:     user_info_result.Data.Name,
		}

		// Generar oauth login
		oauth_login := oauth_logins.OAuthLogin{
			Platform:   "google",
			Email:      user_info_result.Data.Email,
			ExternalID: uuid.New().String(),
		}

		// Crear usuario y email
		controller = saga.SAGA_Controller{
			Steps: []saga.SAGA_Step[any]{
				steps.NewCreateUserStep(
					s.user_repository,
					s.role_repository,
					user,
				),
				oauth_steps.NewCreateOAuthStep(
					s.user_repository,
					s.oauth_repository,
					oauth_login,
				),
			},
		}

		results = controller.Executed(ctx)

		if !controller.Ok() {
			entry.Error("Error al crear usuario y oauth")
			err := errors.New(fmt.Sprintln("Error al crear usuario y oauth: ", controller.Errors()))
			return utils.Responses[entities.SignInResponse]{Err: err}
		}

		oauth_ent = results["entities.OAuthLogin"].Data.(oauth_logins.OAuthLogin)

		user_ent = results["entities.User"].Data.(users.User)

		// Crear login, refresh token y code
		controller_login := saga.SAGA_Controller{
			Steps: []saga.SAGA_Step[any]{
				steps.NewCreateLoginStep(
					s.login_method_repository,
					user_ent.ID,
					oauth_ent.ID,
					"oauth",
				),
			},
			PrevSaga: &controller,
		}

		results_login = controller_login.Executed(ctx)

		if !controller.Ok() {
			entry.Error("Error al crear login")
			err := errors.New(fmt.Sprintln("Error al crear usuario y email: ", controller.Errors()))
			return utils.Responses[entities.SignInResponse]{Err: err}
		}

		entry.Info("Login, Refresh Token and Code created")

		login_ent = results_login["entities.LoginMethod"].Data.(login_methods.LoginMethod)

		refreshs_result := s.createRefreshToken(ctx, login_ent)
		if refreshs_result.Err != nil {
			entry.Error("Error al crear refresh token")
			return utils.Responses[entities.SignInResponse]{
				Err:        refreshs_result.Err,
				StatusCode: 500,
				Success:    false,
			}
		}

		refreshs_ent = refreshs_result.Data
		entry.Info("Refresh Token created")
	} else {
		// Obtenr Oauth
		oauth_ent = oauth[0]
		entry.Info("Oauth found")

		// Obtenr User
		cri = criteria.Criteria{
			Filters: *criteria.NewFilters(
				[]criteria.Filter{
					{
						Field:    "id",
						Operator: criteria.OperatorEqual,
						Value:    oauth_ent.UserID,
					},
				},
			),
		}

		user, err := s.user_repository.Matching(cri)

		if err != nil {
			entry.Error("error matching user")
			return utils.Responses[entities.SignInResponse]{
				Err:        err,
				StatusCode: 500,
				Success:    false,
			}
		}

		if len(user) == 0 {
			entry.Error("user not found")
			return utils.Responses[entities.SignInResponse]{
				Err:        errors.New("user not found"),
				StatusCode: 404,
				Success:    false,
			}
		}

		user_ent = user[0]
		entry.Info("User found")

		// Obtenr Login

		cri = criteria.Criteria{
			Filters: *criteria.NewFilters(
				[]criteria.Filter{
					{
						Field:    "user_id",
						Operator: criteria.OperatorEqual,
						Value:    user_ent.ID,
					},
					{
						Field:    "entity_id",
						Operator: criteria.OperatorEqual,
						Value:    oauth_ent.ID,
					},
					{
						Field:    "entity_type",
						Operator: criteria.OperatorEqual,
						Value:    "oauth",
					},
				},
			),
		}

		login, err := s.login_method_repository.Matching(cri)

		if err != nil {
			entry.Error("error matching login")
			return utils.Responses[entities.SignInResponse]{
				Err:        err,
				StatusCode: 500,
				Success:    false,
			}
		}

		if len(login) == 0 {
			entry.Error("login not found")
			return utils.Responses[entities.SignInResponse]{
				Err:        errors.New("login not found"),
				StatusCode: 404,
				Success:    false,
			}
		}

		login_ent = login[0]

		// Obtenr Refresh Token

		cri = criteria.Criteria{
			Filters: *criteria.NewFilters(
				[]criteria.Filter{
					{
						Field:    "login_method_id",
						Operator: criteria.OperatorEqual,
						Value:    login_ent.ID,
					},
					{
						Field:    "user_id",
						Operator: criteria.OperatorEqual,
						Value:    user_ent.ID,
					},
				},
			),
		}

		refreshs, err := s.refresh_repository.Matching(cri)
		if err != nil {
			entry.Error("error matching refresh token")
			return utils.Responses[entities.SignInResponse]{
				Err:        err,
				StatusCode: 500,
				Success:    false,
			}
		}

		if len(refreshs) == 0 {
			entry.Error("refresh token not found")
			return utils.Responses[entities.SignInResponse]{
				Err:        errors.New("refresh token not found"),
				StatusCode: 404,
				Success:    false,
			}
		}

		refreshs_ent = refreshs[0]
		entry.Info("Refresh Token found")
	}

	result := s.generateTokens(login_ent, refreshs_ent)

	if result.Err != nil {
		return utils.Responses[entities.SignInResponse]{
			StatusCode: 500,
			Errors:     []string{result.Err.Error()},
		}
	}

	s.publishActivationUserEvent(oauth_ent.Email, user_ent.Name)

	// Enviamos el email de binevenida
	return utils.Responses[entities.SignInResponse]{
		Body: entities.SignInResponse{
			JWT:          result.Data.jwt,
			RefreshToken: result.Data.refresh_token,
		},
		StatusCode: 201,
	}

}

func (s OAuthService) createRefreshToken(ctx context.Context, login login_methods.LoginMethod) utils.Either[refreshs.RefreshToken] {

	entry := logger.FromContext(ctx)

	external_id := uuid.New()

	entity := refreshs.RefreshToken{
		UserID:        login.UserID,
		Entity:        domain.Entity{},
		LoginMethodID: login.ID,
		ExternalID:    external_id.String(),
	}

	result := s.refresh_repository.Save(entity)
	if result.Err != nil {
		entry.Error("error saving the refresh token")
		return utils.Either[refreshs.RefreshToken]{Err: result.Err}
	}

	entity.ID = result.Data

	return utils.Either[refreshs.RefreshToken]{
		Data: entity,
	}
}

func (s OAuthService) generateTokens(login login_methods.LoginMethod, refreshToken refreshs.RefreshToken) utils.Either[GenerateTokensFlow] {

	jwt := login.ToJWT(s.jwt_controller)

	refresh_token := refreshToken.ToJWT(s.jwt_controller)

	return utils.Either[GenerateTokensFlow]{Data: GenerateTokensFlow{
		jwt:           jwt,
		refresh_token: refresh_token,
	}}
}

func (s OAuthService) publishActivationUserEvent(email, user_name string) {

	user_event := email_events.UserActivated{
		Email:    email,
		UserName: user_name,
	}

	// Agregar el mensaje a la cola "new-users"
	if err := s.event_bus.Publish([]event.DomainEvent{
		user_event,
	}); err != nil {
		log.Println("Error al publicar el evento new-users")
		log.Println(err)
	}
}
