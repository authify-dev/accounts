package repositories

import (
	"accounts/internal/api/v1/codes/domain/entities"
	"accounts/internal/core/domain/criteria"
	"accounts/internal/utils"
)

// --------------------------------
// DOMAIN
// --------------------------------
// Code Repository
// --------------------------------

type CodeRepository interface {
	Save(role entities.Code) utils.Either[string]
	Search(uuid string) (entities.Code, error)
	SearchAll() ([]entities.Code, error)
	Delete(uuid string) error
	UpdateByFields(uuid string, fields map[string]interface{}) error
	Matching(criteria criteria.Criteria) ([]entities.Code, error)
	View(data []entities.Code)
}
