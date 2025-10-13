package repositories

import (
	"accounts/internal/api/v1/policies/domain/entities"
	"foundation/domain/criteria"
	"foundation/utils"
)

// --------------------------------
// DOMAIN
// --------------------------------
// Policy Repository
// --------------------------------

type PolicyRepository interface {
	Save(policy entities.PolicyEntity) utils.Result[entities.PolicyEntity]
	Search(uuid string) (entities.PolicyEntity, error)
	SearchAll() ([]entities.PolicyEntity, error)
	Delete(uuid string) error
	UpdateByFields(uuid string, fields map[string]interface{}) error
	Matching(criteria criteria.Criteria) ([]entities.PolicyEntity, error)
	View(data []entities.PolicyEntity)
}
