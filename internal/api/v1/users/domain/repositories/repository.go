package repositories

import (
	"accounts/internal/api/v1/users/domain/entities"
	"accounts/internal/core/domain/criteria"
)

// --------------------------------
// DOMAIN
// --------------------------------
// User Repository
// --------------------------------

type UserRepository interface {
	Save(role entities.User) error
	Search(uuid string) (entities.User, error)
	SearchAll() ([]entities.User, error)
	Delete(uuid string) error
	UpdateByFields(uuid string, fields map[string]interface{}) error
	Matching(criteria criteria.Criteria) ([]entities.User, error)
	View(data []entities.User)
}
