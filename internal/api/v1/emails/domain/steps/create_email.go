package steps

import (
	"accounts/internal/api/v1/emails/domain/entities"
	users_entities "accounts/internal/api/v1/users/domain/entities"

	emails "accounts/internal/api/v1/emails/domain/repositories"
	users "accounts/internal/api/v1/users/domain/repositories"
	"accounts/internal/common/logger"
	"accounts/internal/core/domain/criteria"
	"accounts/internal/utils"
	"context"
	"errors"
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
		return utils.Result[any]{Err: err}
	}

	if len(emails) != 0 {
		entry.Error("User already exists")
		return utils.Result[any]{Err: errors.New("user already exists with this email")}
	}

	s.email.UserID = user.ID

	result := s.emails_repo.Save(s.email)
	if result.Err != nil {
		entry.Error("error saving user")
		return utils.Result[any]{Err: result.Err}
	}

	s.email_id = result.Data

	return utils.Result[any]{
		Data: s.email,
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
