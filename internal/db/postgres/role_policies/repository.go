package rolepolicies_pg

import (
	"accounts/internal/api/v1/role_policies/domain/entities"
	"foundation/domain/criteria"
	"foundation/infrastructure/db/cgorm"

	"gorm.io/gorm"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// Role Policies Postgres Repository
// --------------------------------

type RolePoliciesPostgresRepository struct {
	cgorm.PostgresRepository[entities.RolePoliciesEntity, RolePoliciesModel]
}

func NewRolePoliciesPostgresRepository(connection *gorm.DB) *RolePoliciesPostgresRepository {
	return &RolePoliciesPostgresRepository{
		PostgresRepository: cgorm.PostgresRepository[entities.RolePoliciesEntity, RolePoliciesModel]{
			Connection: connection,
		},
	}
}

func (r *RolePoliciesPostgresRepository) Matching(cr criteria.Criteria) ([]entities.RolePoliciesEntity, error) {

	model := &RolePoliciesModel{}

	return r.MatchingLow(cr, model)
}
