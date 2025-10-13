package repositories

import (
	"accounts/internal/context/v1/api_keys/domain/entities"
	"foundation/domain/criteria"
	"foundation/utils"
)

// --------------------------------
// DOMAIN
// --------------------------------
// API Key Repository
// --------------------------------

type APIKeyRepository interface {
	Save(apiKey entities.APIKeyEntity) utils.Result[entities.APIKeyEntity]
	UpdateByFields(uuid string, fields map[string]interface{}) error
	Matching(criteria criteria.Criteria) ([]entities.APIKeyEntity, error)
}
