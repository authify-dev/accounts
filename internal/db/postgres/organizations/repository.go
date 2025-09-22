package organizations_gorm

import (
	"accounts/internal/context/v1/organizations/domain/entities"
	"accounts/internal/core/domain/criteria"
	"accounts/internal/db/postgres"

	"gorm.io/gorm"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// Organizations Postgres Repository
// --------------------------------

type OrganizationsPostgresRepository struct {
	postgres.PostgresRepository[entities.Organization, OrganizationModel]
}

func NewOrganizationsPostgresRepository(connection *gorm.DB) *OrganizationsPostgresRepository {
	return &OrganizationsPostgresRepository{
		PostgresRepository: postgres.PostgresRepository[entities.Organization, OrganizationModel]{
			Connection: connection,
		},
	}
}

func (r *OrganizationsPostgresRepository) Matching(cr criteria.Criteria) ([]entities.Organization, error) {
	model := &OrganizationModel{}

	return r.MatchingLow(cr, model)
}
