package repositories

import (
	"accounts/internal/api/v1/refresh_tokens/domain/entities"
	"accounts/internal/core/domain/criteria"
	"accounts/internal/utils"
)

// --------------------------------
// DOMAIN
// --------------------------------
// RefreshToken Repository
// --------------------------------

type RefreshTokenRepository interface {
	Save(role entities.RefreshToken) utils.Either[string]
	Search(uuid string) (entities.RefreshToken, error)
	SearchAll() ([]entities.RefreshToken, error)
	Delete(uuid string) error
	UpdateByFields(uuid string, fields map[string]interface{}) error
	Matching(criteria criteria.Criteria) ([]entities.RefreshToken, error)
	View(data []entities.RefreshToken)
}
