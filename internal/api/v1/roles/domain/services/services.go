package services

import users "accounts/internal/db/roles"

type RolesService struct {
	repository users.RolesRepository
}

func NewRolesService(repository users.RolesRepository) *RolesService {
	return &RolesService{
		repository: repository,
	}
}
