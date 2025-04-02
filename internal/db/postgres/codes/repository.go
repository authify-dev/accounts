package postgres

import (
	"accounts/internal/api/v1/codes/domain/entities"
	"accounts/internal/core/domain/criteria"
	"accounts/internal/db/postgres"

	"gorm.io/gorm"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// Role Postgres Repository
// --------------------------------

type CodePostgresRepository struct {
	postgres.PostgresRepository[entities.Code, CodeModel]
}

func NewCodePostgresRepository(connection *gorm.DB) *CodePostgresRepository {
	return &CodePostgresRepository{
		PostgresRepository: postgres.PostgresRepository[entities.Code, CodeModel]{
			Connection: connection,
		},
	}
}

func (r *CodePostgresRepository) Matching(cr criteria.Criteria) ([]entities.Code, error) {

	model := &CodeModel{}

	return r.MatchingLow(cr, model)
}
