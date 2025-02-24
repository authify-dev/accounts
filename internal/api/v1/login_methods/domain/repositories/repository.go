package repositories

import (
	"accounts/internal/api/v1/login_methods/domain/entities"
	"accounts/internal/core/domain/criteria"
)

// --------------------------------
// DOMAIN
// --------------------------------
// LoginMethod Repository
// --------------------------------

type LoginMethodRepository interface {
	Save(role entities.LoginMethod) error
	Search(uuid string) (entities.LoginMethod, error)
	SearchAll() ([]entities.LoginMethod, error)
	Delete(uuid string) error
	UpdateByFields(uuid string, fields map[string]interface{}) error
	Matching(criteria criteria.Criteria) ([]entities.LoginMethod, error)
	View(data []entities.LoginMethod)
}
