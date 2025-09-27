package postgres

import (
	"accounts/internal/api/v1/roles/domain/entities"
	"foundation/domain/criteria"
	"foundation/infrastructure/db/cgorm"

	"gorm.io/gorm"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// Role Postgres Repository
// --------------------------------

type RolePostgresRepository struct {
	cgorm.PostgresRepository[entities.Role, RoleModel]
}

func NewRolePostgresRepository(connection *gorm.DB) *RolePostgresRepository {
	return &RolePostgresRepository{
		PostgresRepository: cgorm.PostgresRepository[entities.Role, RoleModel]{
			Connection: connection,
		},
	}
}

func (r *RolePostgresRepository) Matching(cr criteria.Criteria) ([]entities.Role, error) {

	model := &RoleModel{}

	return r.MatchingLow(cr, model)
}
