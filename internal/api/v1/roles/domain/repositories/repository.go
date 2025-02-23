package repositories

import (
	"accounts/internal/api/v1/roles/domain/entities"
	"accounts/internal/core/domain/criteria"
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
	Delete(uuid string) error
	UpdateByFields(uuid string, fields map[string]interface{}) error
	Matching(criteria criteria.Criteria) ([]entities.Role, error)
	View(data []entities.Role)
}
