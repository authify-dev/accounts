package postgres

import (
	"accounts/internal/api/v1/codes/domain/entities"
	"foundation/domain/criteria"
	"foundation/infrastructure/db/cgorm"

	"gorm.io/gorm"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// Role Postgres Repository
// --------------------------------

type CodePostgresRepository struct {
	cgorm.PostgresRepository[entities.Code, CodeModel]
}

func NewCodePostgresRepository(connection *gorm.DB) *CodePostgresRepository {
	return &CodePostgresRepository{
		PostgresRepository: cgorm.PostgresRepository[entities.Code, CodeModel]{
			Connection: connection,
		},
	}
}

func (r *CodePostgresRepository) Matching(cr criteria.Criteria) ([]entities.Code, error) {

	model := &CodeModel{}

	return r.MatchingLow(cr, model)
}
