package apikeys_gorm

import (
	"accounts/internal/context/v1/api_keys/domain/entities"
	"foundation/domain/criteria"
	"foundation/infrastructure/db/cgorm"

	"gorm.io/gorm"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// Role Postgres Repository
// --------------------------------

type APIKeyPostgresRepository struct {
	cgorm.PostgresRepository[entities.APIKeyEntity, APIKeyModel]
}

func NewAPIKeyPostgresRepository(connection *gorm.DB) *APIKeyPostgresRepository {
	return &APIKeyPostgresRepository{
		PostgresRepository: cgorm.PostgresRepository[entities.APIKeyEntity, APIKeyModel]{
			Connection: connection,
		},
	}
}

func (r *APIKeyPostgresRepository) Matching(cr criteria.Criteria) ([]entities.APIKeyEntity, error) {

	model := &APIKeyModel{}

	return r.MatchingLow(cr, model)
}
