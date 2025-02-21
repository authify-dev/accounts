package services

import (
	"accounts/internal/api/v1/users/domain/entities"
	"accounts/internal/common/criteria"
	"fmt"
)

func (u *UsersService) Create(user entities.User, role string) error {

	crit := criteria.Criteria{
		Filters: []criteria.Filter{
			{
				Field:    "name",
				Operator: "=",
				Value:    role,
			},
		},
	}

	roles, err := u.roleRepository.Filters(crit)
	if err != nil {
		return err
	}

	fmt.Println(roles)
	return u.repository.Save(user)
}
