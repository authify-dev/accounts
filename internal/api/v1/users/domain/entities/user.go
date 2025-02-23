package entities

import "accounts/internal/core/domain"

// --------------------------------
// DOMAIN
// --------------------------------
// User Entity
// --------------------------------

// User embebe a Entity, por lo que automáticamente implementa domain.IEntity.
type User struct {
	domain.Entity
	Name     string `json:"name,omitempty"`
	UserName string `json:"user_name,omitempty"`
	RoleID   string `json:"role_id,omitempty"`
	Role     string `json:"role,omitempty"`
}

func (r User) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"id":         r.ID,
		"name":       r.Name,
		"user_name":  r.UserName,
		"created_at": r.CreatedAt,
		"updated_at": r.UpdatedAt,
		"is_removed": r.IsRemoved,
	}
}
