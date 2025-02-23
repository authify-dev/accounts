package services

import (
	roles "accounts/internal/api/v1/roles/domain/repositories"
	users "accounts/internal/api/v1/users/domain/repositories"
)

type UsersService struct {
	repository      users.UserRepository
	role_repository roles.RoleRepository
}

func NewUsersService(
	repository users.UserRepository,
	role_repository roles.RoleRepository,
) *UsersService {
	return &UsersService{
		repository:      repository,
		role_repository: role_repository,
	}
}
