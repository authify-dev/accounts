package repositories

import (
	"accounts/internal/api/v1/emails/domain/entities"
	"accounts/internal/core/domain/criteria"
	"accounts/internal/utils"
)

// --------------------------------
// DOMAIN
// --------------------------------
// Email Repository
// --------------------------------

type EmailRepository interface {
	Save(role entities.Email) utils.Either[string]
	Search(uuid string) (entities.Email, error)
	SearchAll() ([]entities.Email, error)
	Delete(uuid string) error
	UpdateByFields(uuid string, fields map[string]interface{}) error
	Matching(criteria criteria.Criteria) ([]entities.Email, error)
	View(data []entities.Email)
}
