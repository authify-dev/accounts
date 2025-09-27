package organizations_gorm

import (
	"accounts/internal/api/v1/organizations/domain/entities"
	"foundation/domain/criteria"
	"foundation/infrastructure/db/cgorm"

	"gorm.io/gorm"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// Organizations Postgres Repository
// --------------------------------

type OrganizationsPostgresRepository struct {
	cgorm.PostgresRepository[entities.Organization, OrganizationModel]
}

func NewOrganizationsPostgresRepository(connection *gorm.DB) *OrganizationsPostgresRepository {
	return &OrganizationsPostgresRepository{
		PostgresRepository: cgorm.PostgresRepository[entities.Organization, OrganizationModel]{
			Connection: connection,
		},
	}
}

func (r *OrganizationsPostgresRepository) Matching(cr criteria.Criteria) ([]entities.Organization, error) {
	model := &OrganizationModel{}

	return r.MatchingLow(cr, model)
}
