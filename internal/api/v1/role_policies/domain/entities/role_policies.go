package entities

import "accounts/internal/core/domain"

// --------------------------------
// DOMAIN
// --------------------------------
// Role Policies Entity
// --------------------------------

// Role Policies embebe a Entity, por lo que automáticamente implementa domain.IEntity.
type RolePoliciesEntity struct {
	domain.Entity
	RoleID         string `json:"role_id"`
	PolicyID       string `json:"policy_id"`
	OrganizationID string `json:"organization_id"`
}

func (r RolePoliciesEntity) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"id":              r.ID,
		"role_id":         r.RoleID,
		"policy_id":       r.PolicyID,
		"created_at":      r.CreatedAt,
		"updated_at":      r.UpdatedAt,
		"is_removed":      r.IsRemoved,
		"organization_id": r.OrganizationID,
	}
}
