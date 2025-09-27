package entities

import "accounts/internal/core/domain"

// --------------------------------
// DOMAIN
// --------------------------------
// Organization Entity
// --------------------------------

// Organization embebe a Entity, por lo que automáticamente implementa domain.IEntity.
type Organization struct {
	domain.Entity
	Name       string `json:"name"`
	RootUserID string `json:"root_user_id"` // External ID of the root user
}

func (o Organization) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"id":           o.ID,
		"name":         o.Name,
		"root_user_id": o.RootUserID,
		"created_at":   o.CreatedAt,
		"updated_at":   o.UpdatedAt,
		"is_removed":   o.IsRemoved,
	}
}
