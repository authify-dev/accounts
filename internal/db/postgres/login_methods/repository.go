package postgres

import (
	"accounts/internal/api/v1/login_methods/domain/entities"
	"accounts/internal/core/domain/criteria"
	"accounts/internal/db/postgres"

	"gorm.io/gorm"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// Role Postgres Repository
// --------------------------------

type LoginMethodPostgresRepository struct {
	postgres.PostgresRepository[entities.LoginMethod, LoginMethodModel]
}

func NewLoginMethodPostgresRepository(connection *gorm.DB) *LoginMethodPostgresRepository {
	return &LoginMethodPostgresRepository{
		PostgresRepository: postgres.PostgresRepository[entities.LoginMethod, LoginMethodModel]{
			Connection: connection,
		},
	}
}

func (r *LoginMethodPostgresRepository) Matching(cr criteria.Criteria) ([]entities.LoginMethod, error) {

	model := &LoginMethodModel{}

	return r.MatchingLow(cr, model)
}
