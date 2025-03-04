package steps

import (
	roles "accounts/internal/api/v1/roles/domain/repositories"
	"accounts/internal/api/v1/users/domain/entities"
	users "accounts/internal/api/v1/users/domain/repositories"
	"accounts/internal/common/logger"
	"accounts/internal/core/domain/criteria"
	"accounts/internal/utils"
	"context"
	"errors"
)

type CreateUserStep struct {
	user_id   string
	user_repo users.UserRepository
	role_repo roles.RoleRepository
	user      entities.User
}

func NewCreateUserStep(
	user_repo users.UserRepository,
	role_repo roles.RoleRepository,
	user entities.User,
) *CreateUserStep {
	return &CreateUserStep{
		user_repo: user_repo,
		role_repo: role_repo,
		user:      user,
	}
}

func (s *CreateUserStep) Call(ctx context.Context, payload utils.Result[any], allPayloads map[string]utils.Result[any]) utils.Result[any] {

	entry := logger.FromContext(ctx)

	// Verificar ID del role
	criteria := criteria.Criteria{
		Filters: *criteria.NewFilters(
			[]criteria.Filter{
				{
					Field:    "name",
					Value:    s.user.Role,
					Operator: criteria.OperatorEqual,
				},
			},
		),
	}

	roles, err := s.role_repo.Matching(criteria)
	if err != nil {
		entry.Error("error matching role")
		return utils.Result[any]{Err: err}
	}

	if len(roles) == 0 {
		entry.Error("role not found")
		return utils.Result[any]{Err: errors.New("role not found")}
	}

	role_id := roles[0].ID

	// Crear usuario
	s.user.RoleID = role_id

	result := s.user_repo.Save(s.user)

	s.user.ID = result.Data

	if result.Err != nil {
		entry.Error("error saving user")
		return utils.Result[any]{Err: err}
	}

	s.user_id = result.Data

	return utils.Result[any]{
		Data: s.user,
	}
}

func (s *CreateUserStep) Rollback(ctx context.Context) error {
	entry := logger.FromContext(ctx)

	// Implementación de la lógica de negocio
	if s.user_id == "" {
		entry.Error("user_id is empty")
		return nil
	}

	if err := s.user_repo.Delete(s.user_id); err != nil {
		entry.Error("error deleting user")
		return err
	}

	entry.Info("user deleted")
	return nil
}

func (s *CreateUserStep) Produce() string {
	return "entities.User"
}
