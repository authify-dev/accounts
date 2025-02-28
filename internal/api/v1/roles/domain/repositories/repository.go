package repositories

import (
	"accounts/internal/api/v1/roles/domain/entities"
	"accounts/internal/core/domain/criteria"

	"github.com/google/uuid"
)

// --------------------------------
// DOMAIN
// --------------------------------
// Role Repository
// --------------------------------

type RoleRepository interface {
	Save(role entities.Role) error
	Search(uuid string) (entities.Role, error)
	SearchAll() ([]entities.Role, error)
	Delete(uuid uuid.UUID) error
	UpdateByFields(uuid string, fields map[string]interface{}) error
	Matching(criteria criteria.Criteria) ([]entities.Role, error)
	View(data []entities.Role)
}
