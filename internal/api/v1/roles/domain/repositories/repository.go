package repositories

import "accounts/internal/api/v1/roles/domain/entities"

// --------------------------------
// DOMAIN
// --------------------------------
// Role Repository
// --------------------------------

type RoleRepository interface {
	Save(role entities.Role) error
	List() ([]entities.Role, error)
	View(data []entities.Role)
}
