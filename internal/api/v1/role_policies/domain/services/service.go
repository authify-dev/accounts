package services

import (
	policies_repositories "accounts/internal/api/v1/policies/domain/repositories"
	"accounts/internal/api/v1/role_policies/domain/repositories"
	roles_repositories "accounts/internal/api/v1/roles/domain/repositories"
)

type RolePoliciesService struct {
	role_policies_repository repositories.RolePoliciesRepository
	role_repository          roles_repositories.RoleRepository
	policies_repository      policies_repositories.PolicyRepository
}

func NewRolePoliciesService(
	role_policies_repository repositories.RolePoliciesRepository,
	role_repository roles_repositories.RoleRepository,
	policies_repository policies_repositories.PolicyRepository) *RolePoliciesService {
	return &RolePoliciesService{
		role_policies_repository: role_policies_repository,
		role_repository:          role_repository,
		policies_repository:      policies_repository,
	}
}
