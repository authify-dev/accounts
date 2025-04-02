package postgres

import (
	"accounts/internal/api/v1/emails/domain/entities"
	"accounts/internal/core/domain/criteria"
	"accounts/internal/db/postgres"

	"gorm.io/gorm"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// Role Postgres Repository
// --------------------------------

type EmailPostgresRepository struct {
	postgres.PostgresRepository[entities.Email, EmailModel]
}

func NewEmailPostgresRepository(connection *gorm.DB) *EmailPostgresRepository {
	return &EmailPostgresRepository{
		PostgresRepository: postgres.PostgresRepository[entities.Email, EmailModel]{
			Connection: connection,
		},
	}
}

func (r *EmailPostgresRepository) Matching(cr criteria.Criteria) ([]entities.Email, error) {

	model := &EmailModel{}

	return r.MatchingLow(cr, model)
}
