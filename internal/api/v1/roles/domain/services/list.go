package services

import "accounts/internal/api/v1/roles/domain/entities"

func (u *RolesService) List() ([]entities.Role, error) {
	return u.repository.List()
}
