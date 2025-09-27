package repositories

import (
	"accounts/internal/api/v1/emails/domain/entities"
	"foundation/domain/criteria"
	"foundation/utils"
)

// --------------------------------
// DOMAIN
// --------------------------------
// Email Repository
// --------------------------------

type EmailRepository interface {
	Save(role entities.Email) utils.Result[entities.Email]
	Search(uuid string) (entities.Email, error)
	SearchAll() ([]entities.Email, error)
	Delete(uuid string) error
	UpdateByFields(uuid string, fields map[string]interface{}) error
	Matching(criteria criteria.Criteria) ([]entities.Email, error)
	View(data []entities.Email)
}
