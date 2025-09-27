package repositories

import (
	"accounts/internal/api/v1/role_policies/domain/entities"
	"foundation/domain/criteria"
	"foundation/utils"
)

// --------------------------------
// DOMAIN
// --------------------------------
// Role Policies Repository
// --------------------------------

type RolePoliciesRepository interface {
	Save(role_policies entities.RolePoliciesEntity) utils.Result[entities.RolePoliciesEntity]
	Search(uuid string) (entities.RolePoliciesEntity, error)
	SearchAll() ([]entities.RolePoliciesEntity, error)
	Delete(uuid string) error
	UpdateByFields(uuid string, fields map[string]interface{}) error
	Matching(criteria criteria.Criteria) ([]entities.RolePoliciesEntity, error)
	View(data []entities.RolePoliciesEntity)
}
