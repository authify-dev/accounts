package services

import (
	"accounts/internal/api/v1/roles/domain/repositories"
)

type RolesService struct {
	repository repositories.RoleRepository
}

func NewRolesService(repository repositories.RoleRepository) *RolesService {
	return &RolesService{
		repository: repository,
	}
}
