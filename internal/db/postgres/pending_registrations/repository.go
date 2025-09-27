package postgres

import (
	"accounts/internal/api/v1/pending_registrations/domain/entities"
	"foundation/domain/criteria"
	"foundation/infrastructure/db/cgorm"

	"gorm.io/gorm"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// PendingRegistrations Postgres Repository
// --------------------------------

type PendingRegistrationsPostgresRepository struct {
	cgorm.PostgresRepository[entities.PendingRegistration, PendingRegistrationModel]
}

func NewPendingRegistrationsPostgresRepository(connection *gorm.DB) *PendingRegistrationsPostgresRepository {
	return &PendingRegistrationsPostgresRepository{
		PostgresRepository: cgorm.PostgresRepository[entities.PendingRegistration, PendingRegistrationModel]{
			Connection: connection,
		},
	}
}

func (r *PendingRegistrationsPostgresRepository) Matching(cr criteria.Criteria) ([]entities.PendingRegistration, error) {

	model := &PendingRegistrationModel{}

	return r.MatchingLow(cr, model)
}
