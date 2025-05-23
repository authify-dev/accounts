package services

import (
	"accounts/internal/api/v1/emails/domain/entities"
	email_events "accounts/internal/api/v1/emails/domain/events"
	"accounts/internal/common/logger"
	"accounts/internal/core/domain/criteria"
	"accounts/internal/core/domain/event"
	"accounts/internal/utils"
	"context"
	"fmt"
	"log"
	"time"

	code_ents "accounts/internal/api/v1/codes/domain/entities"
	login_ents "accounts/internal/api/v1/login_methods/domain/entities"
)

func (s *EmailsService) ConfirmPassword(
	ctx context.Context,
	entity entities.ConfirmPassword,
) utils.Responses[entities.ResetPasswordResponse] {
	// Logger
	entry := logger.FromContext(ctx)
	entry.Info("Confirm Password")

	// Get email entity
	email := s.getEmail(ctx, entity.Email)
	if email.Err != nil {
		entry.Error("Error al obtener el email")
		return utils.Responses[entities.ResetPasswordResponse]{
			StatusCode: 500,
			Err:        email.Err,
		}
	}

	// Check if the code is valid
	code := s.verifyCode(ctx, email.Data.UserID, "reset_password", entity.Code)
	if code.Err != nil {
		entry.Error("Error al verificar el codigo")
		return utils.Responses[entities.ResetPasswordResponse]{
			StatusCode: 500,
			Err:        code.Err,
		}
	}

	// Get login method
	login := s.getLoginMethod(ctx, email.Data.ID)
	if login.Err != nil {
		entry.Error("Error al obtener el login method")
		return utils.Responses[entities.ResetPasswordResponse]{
			StatusCode: 500,
			Err:        login.Err,
		}
	}

	err := s.login_methods_repository.UpdateByFields(login.Data.ID, map[string]interface{}{
		"is_verify": true,
	})
	if err != nil {
		entry.Error("Error al actualizar el login method")
		return utils.Responses[entities.ResetPasswordResponse]{
			StatusCode: 500,
			Err:        err,
		}
	}

	err = s.codes_repository.UpdateByFields(code.Data.ID, map[string]interface{}{
		"is_removed": true,
	})
	if err != nil {
		entry.Error("Error al actualizar el codigo")
		return utils.Responses[entities.ResetPasswordResponse]{
			StatusCode: 500,
			Err:        err,
		}
	}

	pass_hashed, err := s.password_controller.HashPassword(entity.Password)
	if err != nil {
		entry.Error("Error al hashear la contraseña")
		return utils.Responses[entities.ResetPasswordResponse]{
			StatusCode: 500,
			Err:        err,
		}
	}

	err = s.repository.UpdateByFields(email.Data.ID, map[string]interface{}{
		"password": pass_hashed,
	})

	if err != nil {
		entry.Error("Error al actualizar la contraseña")
		return utils.Responses[entities.ResetPasswordResponse]{
			StatusCode: 500,
			Err:        err,
		}
	}

	s.publishChangedPasswordEvent(email.Data.Email, email.Data.Email)

	return utils.Responses[entities.ResetPasswordResponse]{
		StatusCode: 200,
		Body: entities.ResetPasswordResponse{
			Message: "Password updated",
		},
	}

}

func (s *EmailsService) getLoginMethod(
	ctx context.Context,
	email_id string,
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
			},
		),
	}

	logins, err := s.login_methods_repository.Matching(criteria_login)
	if err != nil {
		entry.Error("Error al obtener el login method")
		return utils.Result[login_ents.LoginMethod]{
			Err: err,
		}
	}

	if len(logins) == 0 {
		entry.Error("Login method not found")
		return utils.Result[login_ents.LoginMethod]{
			Err: fmt.Errorf("login method not found"),
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
			},
		),
	}

	codes, err := s.codes_repository.Matching(criteria_code)
	if err != nil {
		entry.Error("Error al obtener el codigo")
		return utils.Result[code_ents.Code]{
			Err: err,
		}
	}

	if len(codes) == 0 {
		entry.Error("Code not found")
		return utils.Result[code_ents.Code]{
			Err: fmt.Errorf("code not found"),
		}
	}

	code_ent := codes[0]

	if code_ent.Code != code {
		entry.Error("Code not valid")
		return utils.Result[code_ents.Code]{
			Err: fmt.Errorf("code not valid"),
		}
	}

	return utils.Result[code_ents.Code]{
		Data: code_ent,
	}
}

func (s *EmailsService) getEmail(ctx context.Context, email string) utils.Result[entities.Email] {

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
			},
		),
	}

	emails, err := s.repository.Matching(criteria_email)
	if err != nil {
		entry.Error("Error al obtener el email")
		return utils.Result[entities.Email]{
			Err: err,
		}
	}

	if len(emails) == 0 {
		entry.Error("Email not found")
		return utils.Result[entities.Email]{
			Err: fmt.Errorf("email not found"),
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
