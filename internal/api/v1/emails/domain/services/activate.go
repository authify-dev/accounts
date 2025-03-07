package services

import (
	"accounts/internal/api/v1/emails/domain/entities"
	"accounts/internal/common/logger"
	"accounts/internal/core/domain"
	"accounts/internal/core/domain/criteria"
	"accounts/internal/utils"
	"context"
	"fmt"
	"time"

	logins "accounts/internal/api/v1/login_methods/domain/entities"
	refreshs "accounts/internal/api/v1/refresh_tokens/domain/entities"

	"github.com/google/uuid"
)

func (s *EmailsService) Activate(
	ctx context.Context,
	entity entities.Activate,
) utils.Responses[entities.ActivateResponse] {

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

	s.login_methods_repository.UpdateByFields(login.ID, map[string]interface{}{
		"is_verify": true,
	})

	refreshs_result := s.createRefreshToken(ctx, login)

	if refreshs_result.Err != nil {
		return utils.Responses[entities.ActivateResponse]{
			StatusCode: 500,
			Errors:     []string{refreshs_result.Err.Error()},
		}
	}

	result := s.generateTokens(login, refreshs_result.Data)

	if result.Err != nil {
		return utils.Responses[entities.ActivateResponse]{
			StatusCode: 500,
			Errors:     []string{result.Err.Error()},
		}
	}

	return utils.Responses[entities.ActivateResponse]{
		StatusCode: 200,
		Body: entities.ActivateResponse{
			JWT:          result.Data.jwt,
			RefreshToken: result.Data.refresh_token,
		},
	}
}

func (s EmailsService) generateTokens(login logins.LoginMethod, refreshToken refreshs.RefreshToken) utils.Either[GenerateTokensFlow] {

	jwt := login.ToJWT(s.jwt_controller)

	refresh_token := refreshToken.ToJWT(s.jwt_controller)

	return utils.Either[GenerateTokensFlow]{Data: GenerateTokensFlow{
		jwt:           jwt,
		refresh_token: refresh_token,
	}}
}

func (s EmailsService) createRefreshToken(ctx context.Context, login logins.LoginMethod) utils.Either[refreshs.RefreshToken] {

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
		entry.Error("error saving the code")
		return utils.Either[refreshs.RefreshToken]{Err: result.Err}
	}

	entity.ID = result.Data

	return utils.Either[refreshs.RefreshToken]{
		Data: entity,
	}
}
