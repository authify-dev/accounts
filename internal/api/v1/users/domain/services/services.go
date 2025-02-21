package services

import "accounts/internal/db/users"

type UsersService struct {
	repository users.UserRepository
}

func NewUsersService(repository users.UserRepository) *UsersService {
	return &UsersService{
		repository: repository,
	}
}
