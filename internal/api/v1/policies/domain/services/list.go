package services

import (
	"accounts/internal/api/v1/policies/domain/entities"
	"accounts/internal/core/domain/criteria"
	"context"
)

func (s *PoliciesService) List(ctx context.Context, organizationID string) ([]entities.PolicyEntity, error) {

	cri := criteria.Criteria{
		Filters: *criteria.NewFilters(
			[]criteria.Filter{
				{
					Field:    "organization_id",
					Operator: criteria.OperatorEqual,
					Value:    organizationID,
				},
			},
		),
	}

	return s.policies_repository.Matching(cri)
}
