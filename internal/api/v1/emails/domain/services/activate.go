package services

import (
	"accounts/internal/api/v1/emails/domain/entities"
	email_events "accounts/internal/api/v1/emails/domain/events"
	"accounts/internal/common/logger"
	"accounts/internal/core/domain"
	"accounts/internal/core/domain/event"
	"context"
	"fmt"
	"foundation/domain/criteria"
	"foundation/utils"
	"foundation/utils/cerrs"
	"log"
	"net/http"
	"time"

	logins "accounts/internal/api/v1/login_methods/domain/entities"
	refreshs "accounts/internal/api/v1/refresh_tokens/domain/entities"

	"github.com/google/uuid"
)

func (s *EmailsService) Activate(
	ctx context.Context,
	entity entities.Activate,
) utils.Response[entities.ActivateResponse] {

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
		return utils.Response[entities.ActivateResponse]{
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error getting email", "emails.activate.error_getting_email"),
			StatusCode: 500,
		}
	}

	if len(emails) == 0 {
		return utils.Response[entities.ActivateResponse]{
			Error:      cerrs.NewCustomError(http.StatusNotFound, "email not found", "emails.activate.error_email_not_found"),
			StatusCode: 404,
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
		return utils.Response[entities.ActivateResponse]{
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error getting login method", "emails.activate.error_getting_login_method"),
			StatusCode: 500,
		}
	}

	if len(logins) == 0 {
		return utils.Response[entities.ActivateResponse]{
			Error:      cerrs.NewCustomError(http.StatusNotFound, "login not found", "emails.activate.error_login_not_found"),
			StatusCode: 404,
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
		return utils.Response[entities.ActivateResponse]{
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error getting codes", "emails.activate.error_getting_codes"),
			StatusCode: 500,
		}
	}

	if len(codes) == 0 {
		return utils.Response[entities.ActivateResponse]{
			Error:      cerrs.NewCustomError(http.StatusNotFound, "code not found", "emails.activate.error_code_not_found"),
			StatusCode: 404,
		}
	}

	code := codes[0]

	if code.Code != entity.Code {
		return utils.Response[entities.ActivateResponse]{
			Error:      cerrs.NewCustomError(http.StatusBadRequest, "invalid code", "emails.activate.error_invalid_code"),
			StatusCode: 400,
		}
	}

	s.login_methods_repository.UpdateByFields(login.ID.String(), map[string]interface{}{
		"is_verify": true,
	})

	refreshs_result := s.createRefreshToken(ctx, login)

	if refreshs_result.Err != nil {
		return utils.Response[entities.ActivateResponse]{
			Error:      refreshs_result.Err,
			StatusCode: 500,
		}
	}

	s.codes_repository.UpdateByFields(code.ID.String(), map[string]interface{}{
		"is_removed": true,
	})

	result := s.generateTokens(ctx, login, refreshs_result.Data)

	if result.Err != nil {
		return utils.Response[entities.ActivateResponse]{
			Error:      result.Err,
			StatusCode: 500,
		}
	}

	s.publishActivationUserEvent(entity.Email, entity.Email)

	return utils.Response[entities.ActivateResponse]{
		StatusCode: 200,
		Data: entities.ActivateResponse{
			JWT:          result.Data.jwt,
			RefreshToken: result.Data.refresh_token,
		},
	}
}

func (s EmailsService) generateTokens(ctx context.Context, login logins.LoginMethod, refreshToken refreshs.RefreshToken) utils.Result[GenerateTokensFlow] {

	jwt := login.ToJWT(ctx, s.jwt_controller)

	refresh_token := refreshToken.ToJWT(ctx, s.jwt_controller)

	return utils.Result[GenerateTokensFlow]{Data: GenerateTokensFlow{
		jwt:           jwt,
		refresh_token: refresh_token,
	}}
}

func (s EmailsService) createRefreshToken(ctx context.Context, login logins.LoginMethod) utils.Result[refreshs.RefreshToken] {

	entry := logger.FromContext(ctx)

	external_id := uuid.New()

	entity := refreshs.RefreshToken{
		UserID:        login.UserID,
		Entity:        domain.Entity{},
		LoginMethodID: login.ID.String(),
		ExternalID:    external_id.String(),
	}

	result := s.refresh_repository.Save(entity)
	if result.Err != nil {
		entry.Error("error saving the code")
		return utils.Result[refreshs.RefreshToken]{Err: result.Err}
	}

	entity.ID = result.Data.ID

	return utils.Result[refreshs.RefreshToken]{
		Data: entity,
	}
}

func (s EmailsService) publishActivationUserEvent(email, user_name string) {

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
