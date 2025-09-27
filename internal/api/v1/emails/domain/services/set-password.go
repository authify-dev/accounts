package services

import (
	"accounts/internal/api/v1/emails/domain/commands"
	"accounts/internal/api/v1/emails/domain/entities"
	login_ents "accounts/internal/api/v1/login_methods/domain/entities"
	"accounts/internal/common/logger"
	"context"
	"foundation/domain/criteria"
	"foundation/utils"
	"foundation/utils/cerrs"
	"net/http"
)

func (s EmailsService) SetPassword(ctx context.Context, entity commands.SetPassword, jwt string) utils.Response[entities.ActivateResponse] {

	entry := logger.FromContext(ctx)

	entry.Info("Setting password", "email", entity.Email)

	// validar el JWT
	claims, err := s.jwt_controller.ValidateToken(ctx, jwt)
	if err != nil {
		return utils.Response[entities.ActivateResponse]{
			Error:      cerrs.NewCustomError(http.StatusUnauthorized, "invalid token", "emails.set_password.error_invalid_token"),
			StatusCode: 401,
		}
	}

	// obtenemos el user_id del jwt
	user_id := claims["user_id"].(string)

	// obtenemos el id del pending registration
	pending_registration_id := claims["id"].(string)

	// buscamos el pending registration por el id

	cri := criteria.Criteria{
		Filters: *criteria.NewFilters(
			[]criteria.Filter{
				{
					Field:    "id",
					Value:    pending_registration_id,
					Operator: criteria.OperatorEqual,
				},
				{
					Field:    "email",
					Value:    entity.Email,
					Operator: criteria.OperatorEqual,
				},
			},
		),
	}

	pending_registrations, err := s.pending_registrations_repository.Matching(cri)
	if err != nil {
		return utils.Response[entities.ActivateResponse]{
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error getting pending registration", "emails.set_password.error_getting_pending_registration"),
			StatusCode: 500,
		}
	}

	if len(pending_registrations) == 0 {
		return utils.Response[entities.ActivateResponse]{
			Error:      cerrs.NewCustomError(http.StatusNotFound, "pending registration not found", "emails.set_password.error_pending_registration_not_found"),
			StatusCode: 404,
		}
	}

	pending_registration := pending_registrations[0]

	// cxreamos eñ email
	email := entities.Email{
		Email:    entity.Email,
		UserID:   user_id,
		Password: entity.Password,
	}

	emailResult := s.repository.Save(email)
	if emailResult.Err != nil {
		return utils.Response[entities.ActivateResponse]{
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error saving email", "emails.set_password.error_saving_email"),
			StatusCode: 500,
		}
	}

	// creamos un login method vinculando el email y le user
	login_method := login_ents.LoginMethod{
		EntityID:   emailResult.Data.ID.String(),
		EntityType: "email",
		UserID:     user_id,
		IsActive:   true,
		IsVerify:   true,
	}

	login_method_result := s.login_methods_repository.Save(login_method)
	if login_method_result.Err != nil {
		return utils.Response[entities.ActivateResponse]{
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error saving login method", "emails.set_password.error_saving_login_method"),
			StatusCode: 500,
		}
	}

	login_method.ID = login_method_result.Data.ID

	// creamos el refresh token
	refresh_token := s.createRefreshToken(ctx, login_method)
	if refresh_token.Err != nil {
		return utils.Response[entities.ActivateResponse]{
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error creating refresh token", "emails.set_password.error_creating_refresh_token"),
			StatusCode: 500,
		}
	}

	s.codes_repository.UpdateByFields(pending_registration.CodeID, map[string]interface{}{
		"is_removed": true,
	})

	result := s.generateTokens(ctx, login_method, refresh_token.Data)

	if result.Err != nil {
		return utils.Response[entities.ActivateResponse]{
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error generating tokens", "emails.set_password.error_generating_tokens"),
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
