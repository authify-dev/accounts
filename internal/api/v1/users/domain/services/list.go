package services

import (
	"accounts/internal/api/v1/users/domain/entities"
)

func (u *UsersService) List() ([]entities.User, error) {
	return u.repository.SearchAll()
}
