package services

import "accounts/internal/api/v1/roles/domain/entities"

func (u *RolesService) Create(user entities.Role) error {
	return u.repository.Save(user)
}
