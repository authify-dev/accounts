package steps

import (
	logins_entities "accounts/internal/api/v1/login_methods/domain/entities"
	"accounts/internal/core/domain"

	logins "accounts/internal/api/v1/login_methods/domain/repositories"
	"accounts/internal/common/logger"
	"accounts/internal/utils"
	"context"

	"github.com/google/uuid"
)

type CreateLoginStep struct {
	login_id    uuid.UUID
	user_id     uuid.UUID
	entity_id   uuid.UUID
	logins_repo logins.LoginMethodRepository
}

func NewCreateLoginStep(
	logins_repo logins.LoginMethodRepository,
	user_id uuid.UUID,
	entity_id uuid.UUID,
) *CreateLoginStep {
	return &CreateLoginStep{
		logins_repo: logins_repo,
		user_id:     user_id,
		entity_id:   entity_id,
	}
}

func (s *CreateLoginStep) Call(ctx context.Context, payload utils.Result[any], allPayloads map[string]utils.Result[any]) utils.Result[any] {

	entry := logger.FromContext(ctx)

	id := uuid.New()

	login := logins_entities.LoginMethod{
		UserID: s.user_id.String(),
		Entity: domain.Entity{
			ID: id,
		},
		EntityID: s.entity_id.String(),
	}

	err := s.logins_repo.Save(login)
	if err != nil {
		entry.Error("error saving login")
		return utils.Result[any]{Err: err}
	}

	s.login_id = id

	return utils.Result[any]{
		Data: login,
	}
}

func (s *CreateLoginStep) Rollback(ctx context.Context) error {
	entry := logger.FromContext(ctx)

	// Implementación de la lógica de negocio
	if s.login_id == uuid.Nil {
		entry.Error("login_id is empty")
		return nil
	}

	if err := s.logins_repo.Delete(s.login_id); err != nil {
		entry.Error("error deleting user")
		return err
	}

	entry.Info("user deleted")
	return nil
}

func (s *CreateLoginStep) Produce() string {
	return "entities.LoginMethod"
}
