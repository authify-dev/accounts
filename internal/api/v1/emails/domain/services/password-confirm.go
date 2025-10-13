package services

import (
	"accounts/internal/api/v1/emails/domain/commands"
	"accounts/internal/api/v1/emails/domain/entities"
	email_events "accounts/internal/api/v1/emails/domain/events"
	"accounts/internal/common/logger"
	"accounts/internal/core/domain/event"
	"context"
	"foundation/domain/criteria"
	"foundation/domain/customctx"
	"foundation/utils"
	"foundation/utils/cerrs"
	"log"
	"net/http"
	"time"

	code_ents "accounts/internal/api/v1/codes/domain/entities"
	login_ents "accounts/internal/api/v1/login_methods/domain/entities"
)

func (s *EmailsService) ConfirmPassword(
	cc *customctx.CustomContext,
	command commands.ConfirmPassword,
) utils.Response[entities.ResetPasswordResponse] {
	// Logger
	entry := logger.FromContext(cc.Context())
	entry.Info("Confirm Password")

	// Get email entity
	email := s.getEmail(cc.Context(), command.Email, command.OrganizationID)
	if email.Err != nil {
		entry.Error("Error al obtener el email")
		return utils.Response[entities.ResetPasswordResponse]{
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error getting email", "emails.confirm_password.error_getting_email"),
			StatusCode: 500,
		}
	}

	// Check if the code is valid
	code := s.verifyCode(cc.Context(), email.Data.UserID, "reset_password", command.Code, command.OrganizationID)
	if code.Err != nil {
		entry.Error("Error al verificar el codigo")
		return utils.Response[entities.ResetPasswordResponse]{
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error verifying code", "emails.confirm_password.error_verifying_code"),
			StatusCode: 500,
		}
	}

	// Get login method
	login := s.getLoginMethod(cc.Context(), email.Data.ID.String(), command.OrganizationID)
	if login.Err != nil {
		entry.Error("Error al obtener el login method")
		return utils.Response[entities.ResetPasswordResponse]{
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error getting login method", "emails.confirm_password.error_getting_login_method"),
			StatusCode: 500,
		}
	}

	err := s.login_methods_repository.UpdateByFields(login.Data.ID.String(), map[string]interface{}{
		"is_verify": true,
	})
	if err != nil {
		entry.Error("Error al actualizar el login method")
		return utils.Response[entities.ResetPasswordResponse]{
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error updating login method", "emails.confirm_password.error_updating_login_method"),
			StatusCode: 500,
		}
	}

	err = s.codes_repository.UpdateByFields(code.Data.ID.String(), map[string]interface{}{
		"is_removed": true,
	})
	if err != nil {
		entry.Error("Error al actualizar el codigo")
		return utils.Response[entities.ResetPasswordResponse]{
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error updating code", "emails.confirm_password.error_updating_code"),
			StatusCode: 500,
		}
	}

	pass_hashed, err := s.password_controller.HashPassword(command.Password)
	if err != nil {
		entry.Error("Error al hashear la contraseña")
		return utils.Response[entities.ResetPasswordResponse]{
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error hashing password", "emails.confirm_password.error_hashing_password"),
			StatusCode: 500,
		}
	}

	err = s.repository.UpdateByFields(email.Data.ID.String(), map[string]interface{}{
		"password": pass_hashed,
	})

	if err != nil {
		entry.Error("Error al actualizar la contraseña")
		return utils.Response[entities.ResetPasswordResponse]{
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error updating password", "emails.confirm_password.error_updating_password"),
			StatusCode: 500,
		}
	}

	s.publishChangedPasswordEvent(command.Email, email.Data.Email)

	return utils.Response[entities.ResetPasswordResponse]{
		StatusCode: 200,
		Data: entities.ResetPasswordResponse{
			Message: "Password updated",
		},
	}

}

func (s *EmailsService) getLoginMethod(
	ctx context.Context,
	email_id string,
	organization_id string,
) utils.Result[login_ents.LoginMethod] {
	// Logger
	entry := logger.FromContext(ctx)
	entry.Info("Get Login Method")

	criteria_login := criteria.Criteria{
		Filters: *criteria.NewFilters(
			[]criteria.Filter{
				{
					Field:    "entity_id",
					Value:    email_id,
					Operator: criteria.OperatorEqual,
				},
				{
					Field:    "entity_type",
					Value:    "email",
					Operator: criteria.OperatorEqual,
				},
				{
					Field:    "organization_id",
					Value:    organization_id,
					Operator: criteria.OperatorEqual,
				},
			},
		),
	}

	logins, err := s.login_methods_repository.Matching(criteria_login)
	if err != nil {
		entry.Error("Error al obtener el login method")
		return utils.Result[login_ents.LoginMethod]{
			Err: cerrs.NewCustomError(http.StatusInternalServerError, "error getting login method", "emails.confirm_password.error_getting_login_method"),
		}
	}

	if len(logins) == 0 {
		entry.Error("Login method not found")
		return utils.Result[login_ents.LoginMethod]{
			Err: cerrs.NewCustomError(http.StatusNotFound, "login method not found", "emails.confirm_password.error_login_method_not_found"),
		}
	}

	login := logins[0]

	return utils.Result[login_ents.LoginMethod]{
		Data: login,
	}
}

func (s *EmailsService) verifyCode(
	ctx context.Context,
	user_id string,
	type_code string,
	code string,
	organization_id string,
) utils.Result[code_ents.Code] {
	// Logger
	entry := logger.FromContext(ctx)
	entry.Info("Verify Code")

	criteria_code := criteria.Criteria{
		Filters: *criteria.NewFilters(
			[]criteria.Filter{
				{
					Field:    "user_id",
					Value:    user_id,
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
				{
					Field:    "type",
					Value:    type_code,
					Operator: criteria.OperatorEqual,
				},
				{
					Field:    "organization_id",
					Value:    organization_id,
					Operator: criteria.OperatorEqual,
				},
			},
		),
	}

	codes, err := s.codes_repository.Matching(criteria_code)
	if err != nil {
		entry.Error("Error al obtener el codigo")
		return utils.Result[code_ents.Code]{
			Err: cerrs.NewCustomError(http.StatusInternalServerError, "error getting code", "emails.confirm_password.error_getting_code"),
		}
	}

	if len(codes) == 0 {
		entry.Error("Code not found")
		return utils.Result[code_ents.Code]{
			Err: cerrs.NewCustomError(http.StatusNotFound, "code not found", "emails.confirm_password.error_code_not_found"),
		}
	}

	code_ent := codes[0]

	if code_ent.Code != code {
		entry.Error("Code not valid")
		return utils.Result[code_ents.Code]{
			Err: cerrs.NewCustomError(http.StatusUnauthorized, "code not valid", "emails.confirm_password.error_code_not_valid"),
		}
	}

	return utils.Result[code_ents.Code]{
		Data: code_ent,
	}
}

func (s *EmailsService) getEmail(ctx context.Context, email string, organization_id string) utils.Result[entities.Email] {

	// Logger
	entry := logger.FromContext(ctx)
	entry.Info("Get Email")

	// Get email entity
	criteria_email := criteria.Criteria{
		Filters: *criteria.NewFilters(
			[]criteria.Filter{
				{
					Field:    "email",
					Value:    email,
					Operator: criteria.OperatorEqual,
				},
				{
					Field:    "organization_id",
					Value:    organization_id,
					Operator: criteria.OperatorEqual,
				},
			},
		),
	}

	emails, err := s.repository.Matching(criteria_email)
	if err != nil {
		entry.Error("Error al obtener el email")
		return utils.Result[entities.Email]{
			Err: cerrs.NewCustomError(http.StatusInternalServerError, "error getting email", "emails.confirm_password.error_getting_email"),
		}
	}

	if len(emails) == 0 {
		entry.Error("Email not found")
		return utils.Result[entities.Email]{
			Err: cerrs.NewCustomError(http.StatusNotFound, "email not found", "emails.confirm_password.error_email_not_found"),
		}
	}

	emailEntity := emails[0]

	return utils.Result[entities.Email]{
		Data: emailEntity,
	}
}

func (s EmailsService) publishChangedPasswordEvent(email string, user_name string) {

	user_event := email_events.ChangedPassword{
		Email:    email,
		UserName: user_name,
	}

	// Agregar el mensaje a la cola "new-users"
	if err := s.event_bus.Publish([]event.DomainEvent{
		user_event,
	}); err != nil {
		log.Println("Error al publicar el evento changed_password")
		log.Println(err)
	}
}
