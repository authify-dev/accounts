package services

import (
	"accounts/internal/api/v1/roles/domain/entities"
	"foundation/domain/criteria"
)

func (u *RolesService) List(organizationID string) ([]entities.Role, error) {
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
	return u.repository.Matching(cri)
}
