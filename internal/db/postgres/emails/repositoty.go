package postgres

import (
	"accounts/internal/api/v1/emails/domain/entities"
	"foundation/domain/criteria"
	"foundation/infrastructure/db/cgorm"

	"gorm.io/gorm"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// Role Postgres Repository
// --------------------------------

type EmailPostgresRepository struct {
	cgorm.PostgresRepository[entities.Email, EmailModel]
}

func NewEmailPostgresRepository(connection *gorm.DB) *EmailPostgresRepository {
	return &EmailPostgresRepository{
		PostgresRepository: cgorm.PostgresRepository[entities.Email, EmailModel]{
			Connection: connection,
		},
	}
}

func (r *EmailPostgresRepository) Matching(cr criteria.Criteria) ([]entities.Email, error) {

	model := &EmailModel{}

	return r.MatchingLow(cr, model)
}
