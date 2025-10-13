package steps

import (
	"accounts/internal/api/v1/emails/domain/entities"
	users_entities "accounts/internal/api/v1/users/domain/entities"
	"context"
	"net/http"

	emails "accounts/internal/api/v1/emails/domain/repositories"
	users "accounts/internal/api/v1/users/domain/repositories"
	"accounts/internal/common/logger"
	"foundation/domain/criteria"
	"foundation/utils"
	"foundation/utils/cerrs"
)

type CreateEmailStep struct {
	email_id    string
	user_repo   users.UserRepository
	emails_repo emails.EmailRepository
	email       entities.Email
}

func NewCreateEmailStep(
	user_repo users.UserRepository,
	emails_repo emails.EmailRepository,
	email entities.Email,
) *CreateEmailStep {
	return &CreateEmailStep{
		user_repo:   user_repo,
		emails_repo: emails_repo,
		email:       email,
	}
}

func (s *CreateEmailStep) Call(ctx context.Context, payload utils.Result[any], allPayloads map[string]utils.Result[any]) utils.Result[any] {

	entry := logger.FromContext(ctx)

	user := payload.Data.(users_entities.User)

	// Verificar ID del role
	criteria := criteria.Criteria{
		Filters: *criteria.NewFilters(
			[]criteria.Filter{
				{
					Field:    "email",
					Value:    s.email.Email,
					Operator: criteria.OperatorEqual,
				},
			},
		),
	}

	emails, err := s.emails_repo.Matching(criteria)
	if err != nil {
		entry.Error("error matching role")
		return utils.Result[any]{Err: cerrs.NewCustomError(http.StatusInternalServerError, "error matching role", "emails.create_email.error_matching_role")}
	}

	if len(emails) != 0 {
		entry.Error("User already exists")
		return utils.Result[any]{Err: cerrs.NewCustomError(http.StatusBadRequest, "user already exists with this email", "emails.create_email.user_already_exists_with_this_email")}
	}

	s.email.UserID = user.ID.String()

	result := s.emails_repo.Save(s.email)
	if result.Err != nil {
		entry.Error("error saving user")
		return utils.Result[any]{Err: cerrs.NewCustomError(http.StatusInternalServerError, "error saving user", "emails.create_email.error_saving_user")}
	}

	s.email_id = result.Data.ID.String()

	return utils.Result[any]{
		Data: result.Data,
	}
}

func (s *CreateEmailStep) Rollback(ctx context.Context) error {
	entry := logger.FromContext(ctx)

	// Implementación de la lógica de negocio
	if s.email_id == "" {
		entry.Error("email_id is empty")
		return nil
	}

	if err := s.emails_repo.Delete(s.email_id); err != nil {
		entry.Error("error deleting user")
		return err
	}

	entry.Info("user deleted")
	return nil
}

func (s *CreateEmailStep) Produce() string {
	return "entities.Email"
}
