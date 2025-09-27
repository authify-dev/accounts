package repositories

import (
	"accounts/internal/api/v1/refresh_tokens/domain/entities"
	"foundation/domain/criteria"
	"foundation/utils"
)

// --------------------------------
// DOMAIN
// --------------------------------
// RefreshToken Repository
// --------------------------------

type RefreshTokenRepository interface {
	Save(role entities.RefreshToken) utils.Result[entities.RefreshToken]
	Search(uuid string) (entities.RefreshToken, error)
	SearchAll() ([]entities.RefreshToken, error)
	Delete(uuid string) error
	UpdateByFields(uuid string, fields map[string]interface{}) error
	Matching(criteria criteria.Criteria) ([]entities.RefreshToken, error)
	View(data []entities.RefreshToken)
}
