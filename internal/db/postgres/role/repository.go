package postgres

import (
	"accounts/internal/api/v1/roles/domain/entities"
	"accounts/internal/core/domain/criteria"
	"accounts/internal/db/postgres"

	"gorm.io/gorm"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// Role Postgres Repository
// --------------------------------

type RolePostgresRepository struct {
	postgres.PostgresRepository[entities.Role, RoleModel]
}

func NewRolePostgresRepository(connection *gorm.DB) *RolePostgresRepository {
	return &RolePostgresRepository{
		PostgresRepository: postgres.PostgresRepository[entities.Role, RoleModel]{
			Connection: connection,
		},
	}
}

func (r *RolePostgresRepository) Matching(cr criteria.Criteria) ([]entities.Role, error) {

	model := &RoleModel{}

	return r.MatchingLow(cr, model)
}
