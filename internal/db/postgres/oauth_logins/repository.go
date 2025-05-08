package postgres

import (
	"accounts/internal/api/v1/oauth_logins/domain/entities"
	"accounts/internal/core/domain/criteria"
	"accounts/internal/db/postgres"

	"gorm.io/gorm"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// OAuth Postgres Repository
// --------------------------------

type OAuthLoginPostgresRepository struct {
	postgres.PostgresRepository[entities.OAuthLogin, OAuthLoginModel]
}

func NewOAuthLoginPostgresRepository(connection *gorm.DB) *OAuthLoginPostgresRepository {
	return &OAuthLoginPostgresRepository{
		PostgresRepository: postgres.PostgresRepository[entities.OAuthLogin, OAuthLoginModel]{
			Connection: connection,
		},
	}
}

func (r *OAuthLoginPostgresRepository) Matching(cr criteria.Criteria) ([]entities.OAuthLogin, error) {
	model := &OAuthLoginModel{}

	return r.MatchingLow(cr, model)
}
