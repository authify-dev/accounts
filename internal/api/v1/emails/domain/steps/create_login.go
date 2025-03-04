package steps

import (
	logins_entities "accounts/internal/api/v1/login_methods/domain/entities"
	"accounts/internal/core/domain"

	logins "accounts/internal/api/v1/login_methods/domain/repositories"
	"accounts/internal/common/logger"
	"accounts/internal/utils"
	"context"
)

type CreateLoginStep struct {
	login_id    string
	user_id     string
	entity_id   string
	logins_repo logins.LoginMethodRepository
}

func NewCreateLoginStep(
	logins_repo logins.LoginMethodRepository,
	user_id string,
	entity_id string,
) *CreateLoginStep {
	return &CreateLoginStep{
		logins_repo: logins_repo,
		user_id:     user_id,
		entity_id:   entity_id,
	}
}

func (s *CreateLoginStep) Call(ctx context.Context, payload utils.Result[any], allPayloads map[string]utils.Result[any]) utils.Result[any] {

	entry := logger.FromContext(ctx)

	login := logins_entities.LoginMethod{
		UserID:   s.user_id,
		Entity:   domain.Entity{},
		EntityID: s.entity_id,
	}

	result := s.logins_repo.Save(login)
	if result.Err != nil {
		entry.Error("error saving login")
		return utils.Result[any]{Err: result.Err}
	}

	s.login_id = result.Data
	login.ID = result.Data

	return utils.Result[any]{
		Data: login,
	}
}

func (s *CreateLoginStep) Rollback(ctx context.Context) error {
	entry := logger.FromContext(ctx)

	// Implementación de la lógica de negocio
	if s.login_id == "" {
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
