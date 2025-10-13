package services

import (
	"accounts/internal/api/v1/policies/domain/repositories"
)

type PoliciesService struct {
	policies_repository repositories.PolicyRepository
}

func NewPoliciesService(policies_repository repositories.PolicyRepository) *PoliciesService {
	return &PoliciesService{policies_repository: policies_repository}
}
