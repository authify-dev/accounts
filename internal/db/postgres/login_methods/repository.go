package postgres

import (
	"accounts/internal/api/v1/login_methods/domain/entities"
	"foundation/domain/criteria"
	"foundation/infrastructure/db/cgorm"

	"gorm.io/gorm"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// Role Postgres Repository
// --------------------------------

type LoginMethodPostgresRepository struct {
	cgorm.PostgresRepository[entities.LoginMethod, LoginMethodModel]
}

func NewLoginMethodPostgresRepository(connection *gorm.DB) *LoginMethodPostgresRepository {
	return &LoginMethodPostgresRepository{
		PostgresRepository: cgorm.PostgresRepository[entities.LoginMethod, LoginMethodModel]{
			Connection: connection,
		},
	}
}

func (r *LoginMethodPostgresRepository) Matching(cr criteria.Criteria) ([]entities.LoginMethod, error) {

	model := &LoginMethodModel{}

	return r.MatchingLow(cr, model)
}
