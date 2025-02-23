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
	List() ([]entities.Role, error)
	View(data []entities.Role)
	Matching(criteria criteria.Criteria) ([]entities.Role, error)
}
