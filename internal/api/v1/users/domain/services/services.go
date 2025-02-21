package services

import (
	"accounts/internal/db/roles"
	"accounts/internal/db/users"
)

type UsersService struct {
	repository     users.UserRepository
	roleRepository roles.RolesRepository
}

func NewUsersService(
	repository users.UserRepository,
	roleRepository roles.RolesRepository,
) *UsersService {
	return &UsersService{
		repository:     repository,
		roleRepository: roleRepository,
	}
}
