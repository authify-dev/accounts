package postgres

import (
	"accounts/internal/api/v1/users/domain/entities"
	"accounts/internal/core/domain/criteria"
	"accounts/internal/db/postgres"

	"gorm.io/gorm"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// User Postgres Repository
// --------------------------------

type UserPostgresRepository struct {
	postgres.PostgresRepository[entities.User, UserModel]
}

func NewUserPostgresRepository(connection *gorm.DB) *UserPostgresRepository {
	return &UserPostgresRepository{
		PostgresRepository: postgres.PostgresRepository[entities.User, UserModel]{
			Connection: connection,
		},
	}
}

func (r *UserPostgresRepository) Matching(cr criteria.Criteria) ([]entities.User, error) {

	model := &UserModel{}

	return r.MatchingLow(cr, model)
}
