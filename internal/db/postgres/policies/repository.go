package policies_pg

import (
	"accounts/internal/api/v1/policies/domain/entities"
	"foundation/domain/criteria"
	"foundation/infrastructure/db/cgorm"

	"gorm.io/gorm"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// Policies Postgres Repository
// --------------------------------

type PoliciesPostgresRepository struct {
	cgorm.PostgresRepository[entities.PolicyEntity, PolicyModel]
}

func NewPoliciesPostgresRepository(connection *gorm.DB) *PoliciesPostgresRepository {
	return &PoliciesPostgresRepository{
		PostgresRepository: cgorm.PostgresRepository[entities.PolicyEntity, PolicyModel]{
			Connection: connection,
		},
	}
}

func (r *PoliciesPostgresRepository) Matching(cr criteria.Criteria) ([]entities.PolicyEntity, error) {

	model := &PolicyModel{}

	return r.MatchingLow(cr, model)
}
