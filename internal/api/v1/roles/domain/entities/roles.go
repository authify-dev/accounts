package entities

import "accounts/internal/core/domain"

// --------------------------------
// DOMAIN
// --------------------------------
// Role Entity
// --------------------------------

// Role embebe a Entity, por lo que automáticamente implementa domain.IEntity.
type Role struct {
	domain.Entity
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

func (r Role) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"id":          r.ID,
		"name":        r.Name,
		"description": r.Description,
		"created_at":  r.CreatedAt,
		"updated_at":  r.UpdatedAt,
		"is_removed":  r.IsRemoved,
	}
}
