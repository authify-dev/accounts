package repositories

import (
	"accounts/internal/api/v1/pending_registrations/domain/entities"
	"foundation/domain/criteria"
	"foundation/utils"
)

// --------------------------------
// DOMAIN
// --------------------------------
// Pending Registrations Repository
// --------------------------------

type PendingRegistrationsRepository interface {
	Save(role entities.PendingRegistration) utils.Result[entities.PendingRegistration]
	Search(uuid string) (entities.PendingRegistration, error)
	SearchAll() ([]entities.PendingRegistration, error)
	Delete(uuid string) error
	UpdateByFields(uuid string, fields map[string]interface{}) error
	Matching(criteria criteria.Criteria) ([]entities.PendingRegistration, error)
	View(data []entities.PendingRegistration)
}
