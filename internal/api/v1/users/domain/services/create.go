package services

import "accounts/internal/api/v1/users/domain/entities"

func (u *UsersService) Create(user entities.User) error {
	return u.repository.Save(user)
}
