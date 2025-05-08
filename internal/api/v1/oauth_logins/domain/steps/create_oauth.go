package steps

import (
	"accounts/internal/api/v1/oauth_logins/domain/entities"
	oauth "accounts/internal/api/v1/oauth_logins/domain/repositories"
	users_entities "accounts/internal/api/v1/users/domain/entities"
	users "accounts/internal/api/v1/users/domain/repositories"
	"fmt"

	"accounts/internal/common/logger"
	"accounts/internal/core/domain/criteria"
	"accounts/internal/utils"
	"context"
)

type CreateOAuthStep struct {
	oauth_id   string
	user_repo  users.UserRepository
	oauth_repo oauth.OAuthLoginRepository
	oauth      entities.OAuthLogin
}

func NewCreateOAuthStep(
	user_repo users.UserRepository,
	oauth_repo oauth.OAuthLoginRepository,
	oauth entities.OAuthLogin,
) *CreateOAuthStep {
	return &CreateOAuthStep{
		user_repo:  user_repo,
		oauth_repo: oauth_repo,
		oauth:      oauth,
	}
}

func (s *CreateOAuthStep) Call(ctx context.Context, payload utils.Result[any], allPayloads map[string]utils.Result[any]) utils.Result[any] {

	entry := logger.FromContext(ctx)

	// Verificar el ID del user
	user := payload.Data.(users_entities.User)

	// Verificar ID del role
	criteria := criteria.Criteria{
		Filters: *criteria.NewFilters(
			[]criteria.Filter{
				{
					Field:    "email",
					Value:    s.oauth.Email,
					Operator: criteria.OperatorEqual,
				},
			},
		),
	}

	oauths, err := s.oauth_repo.Matching(criteria)

	if err != nil {
		entry.Error("error matching oauth")
		return utils.Result[any]{Err: err}
	}

	if len(oauths) != 0 {
		entry.Error("oauth already exists")
		return utils.Result[any]{Err: fmt.Errorf("oauth already exists")}
	}

	s.oauth.UserID = user.ID

	result := s.oauth_repo.Save(s.oauth)
	if result.Err != nil {
		entry.Error("error saving oauth")
		return utils.Result[any]{Err: result.Err}
	}

	s.oauth_id = result.Data
	s.oauth.ID = s.oauth_id

	return utils.Result[any]{
		Data: s.oauth,
	}
}

func (s *CreateOAuthStep) Rollback(ctx context.Context) error {
	entry := logger.FromContext(ctx)

	// Implementación de la lógica de negocio
	if s.oauth_id == "" {
		entry.Error("oauth_id is empty")
		return nil
	}

	if err := s.user_repo.Delete(s.oauth_id); err != nil {
		entry.Error("error deleting oauth")
		return err
	}

	entry.Info("oauth deleted")
	return nil
}

func (s *CreateOAuthStep) Produce() string {
	return "entities.OAuthLogin"
}
