package postgres

import (
	"accounts/internal/api/v1/users/domain/entities"
	"foundation/domain/criteria"
	"foundation/infrastructure/db/cgorm"

	"gorm.io/gorm"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// User Postgres Repository
// --------------------------------

type UserPostgresRepository struct {
	cgorm.PostgresRepository[entities.User, UserModel]
}

func NewUserPostgresRepository(connection *gorm.DB) *UserPostgresRepository {
	return &UserPostgresRepository{
		PostgresRepository: cgorm.PostgresRepository[entities.User, UserModel]{
			Connection: connection,
		},
	}
}

func (r *UserPostgresRepository) Matching(cr criteria.Criteria) ([]entities.User, error) {

	model := &UserModel{}

	return r.MatchingLow(cr, model)
}
