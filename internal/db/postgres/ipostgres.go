package postgres

import (
	"accounts/internal/core/domain"
	"accounts/internal/core/domain/criteria"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// PostgresRepository
// --------------------------------

type IPostgresRepository[E domain.IEntity, M domain.IModel] interface {
	View(data []E)
	Save(role E) error
	SearchAll() ([]E, error)
	MatchingLow(cr criteria.Criteria, model *M) ([]E, error)
	Delete(uuid string) error
	Search(uuid string) (E, error)
	UpdateByFields(uuid string, fields map[string]interface{}) error
}
