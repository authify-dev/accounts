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
	connection *gorm.DB
}

func NewUserPostgresRepository(connection *gorm.DB) *UserPostgresRepository {
	return &UserPostgresRepository{connection: connection}
}

func (r *UserPostgresRepository) Matching(criteria criteria.Criteria) ([]entities.User, error) {

	return r.List()
}
