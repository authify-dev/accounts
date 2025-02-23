package postgres

import (
	"accounts/internal/api/v1/refresh_tokens/domain/entities"
	"accounts/internal/core/domain/criteria"
	"accounts/internal/db/postgres"

	"gorm.io/gorm"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// Role Postgres Repository
// --------------------------------

type RefreshTokenPostgresRepository struct {
	postgres.PostgresRepository[entities.RefreshToken, RefreshTokenModel]
}

func NewRefreshTokenPostgresRepository(connection *gorm.DB) *RefreshTokenPostgresRepository {
	return &RefreshTokenPostgresRepository{
		PostgresRepository: postgres.PostgresRepository[entities.RefreshToken, RefreshTokenModel]{
			Connection: connection,
		},
	}
}

func (r *RefreshTokenPostgresRepository) Matching(cr criteria.Criteria) ([]entities.RefreshToken, error) {

	model := &RefreshTokenModel{}

	return r.MatchingLow(cr, model)
}
