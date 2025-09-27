package entities

import "accounts/internal/core/domain"

// --------------------------------
// DOMAIN
// --------------------------------
// APIKey Entity
// --------------------------------

// APIKey embebe a Entity, por lo que automáticamente implementa domain.IEntity.
type APIKeyEntity struct {
	domain.Entity
	Name           string `json:"name"`
	Description    string `json:"description"`
	OrganizationID string `json:"organization_id"`
	Key            string `json:"key"`
}

func (o APIKeyEntity) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"id":              o.ID,
		"name":            o.Name,
		"description":     o.Description,
		"organization_id": o.OrganizationID,
		"key":             o.Key,
		"created_at":      o.CreatedAt,
		"updated_at":      o.UpdatedAt,
		"is_removed":      o.IsRemoved,
	}
}
