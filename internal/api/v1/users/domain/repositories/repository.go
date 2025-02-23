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
	List() ([]entities.User, error)
	View(data []entities.User)
	Matching(criteria criteria.Criteria) ([]entities.User, error)
}
