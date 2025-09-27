package repositories

import (
	"accounts/internal/api/v1/codes/domain/entities"
	"foundation/domain/criteria"
	"foundation/utils"
)

// --------------------------------
// DOMAIN
// --------------------------------
// Code Repository
// --------------------------------

type CodeRepository interface {
	Save(role entities.Code) utils.Result[entities.Code]
	Search(uuid string) (entities.Code, error)
	SearchAll() ([]entities.Code, error)
	Delete(uuid string) error
	UpdateByFields(uuid string, fields map[string]interface{}) error
	Matching(criteria criteria.Criteria) ([]entities.Code, error)
	View(data []entities.Code)
}
