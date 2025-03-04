package services

import (
	"accounts/internal/api/v1/users/domain/entities"
	"accounts/internal/core/domain/criteria"
	"errors"
)

func (u *UsersService) Create(user entities.User) error {

	cri := criteria.Criteria{
		Filters: *criteria.NewFilters(
			[]criteria.Filter{
				{
					Field:    "name",
					Operator: criteria.OperatorEqual,
					Value:    user.Role,
				},
			},
		),
	}

	roles, err := u.role_repository.Matching(cri)
	if err != nil {
		return err
	}

	if len(roles) == 0 {
		return errors.New("Role not found")
	}

	role := roles[0]

	user.RoleID = role.GetID()

	return u.repository.Save(user).Err
}
