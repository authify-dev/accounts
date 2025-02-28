package repositories

import (
	"accounts/internal/api/v1/codes/domain/entities"
	"accounts/internal/core/domain/criteria"

	"github.com/google/uuid"
)

// --------------------------------
// DOMAIN
// --------------------------------
// Code Repository
// --------------------------------

type CodeRepository interface {
	Save(role entities.Code) error
	Search(uuid string) (entities.Code, error)
	SearchAll() ([]entities.Code, error)
	Delete(uuid uuid.UUID) error
	UpdateByFields(uuid string, fields map[string]interface{}) error
	Matching(criteria criteria.Criteria) ([]entities.Code, error)
	View(data []entities.Code)
}
