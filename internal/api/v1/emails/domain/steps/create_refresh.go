package steps

import (
	logins_entities "accounts/internal/api/v1/login_methods/domain/entities"
	refresh_tokens_entities "accounts/internal/api/v1/refresh_tokens/domain/entities"

	"accounts/internal/core/domain"

	refresh "accounts/internal/api/v1/refresh_tokens/domain/repositories"
	"accounts/internal/common/logger"
	"accounts/internal/utils"
	"context"

	"github.com/google/uuid"
)

type CreateRefreshTokenStep struct {
	refresh_token_id string
	user_id          string
	refresh_repo     refresh.RefreshTokenRepository
}

func NewCreateRefreshTokenStep(
	refresh_repo refresh.RefreshTokenRepository,
	user_id string,
) *CreateRefreshTokenStep {
	return &CreateRefreshTokenStep{
		refresh_repo: refresh_repo,
		user_id:      user_id,
	}
}

func (s *CreateRefreshTokenStep) Call(ctx context.Context, payload utils.Result[any], allPayloads map[string]utils.Result[any]) utils.Result[any] {

	entry := logger.FromContext(ctx)

	login := payload.Data.(logins_entities.LoginMethod)

	external_id := uuid.New()

	entity := refresh_tokens_entities.RefreshToken{
		UserID:        s.user_id,
		Entity:        domain.Entity{},
		LoginMethodID: login.ID,
		ExternalID:    external_id.String(),
	}

	result := s.refresh_repo.Save(entity)
	if result.Err != nil {
		entry.Error("error saving the code")
		return utils.Result[any]{Err: result.Err}
	}

	s.refresh_token_id = result.Data

	return utils.Result[any]{
		Data: entity,
	}
}

func (s *CreateRefreshTokenStep) Rollback(ctx context.Context) error {
	entry := logger.FromContext(ctx)

	// Implementación de la lógica de negocio
	if s.refresh_token_id == "" {
		entry.Error("refresh_token_id is empty")
		return nil
	}

	if err := s.refresh_repo.Delete(s.refresh_token_id); err != nil {
		entry.Error("error deleting user")
		return err
	}

	entry.Info("user deleted")
	return nil
}

func (s *CreateRefreshTokenStep) Produce() string {
	return "entities.RefreshToken"
}
