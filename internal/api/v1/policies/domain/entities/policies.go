package entities

import (
	"accounts/internal/core/domain"
	"encoding/json"
)

// --------------------------------
// DOMAIN
// --------------------------------
// Policy Entity
// --------------------------------

// Role embebe a Entity, por lo que automáticamente implementa domain.IEntity.
type PolicyEntity struct {
	domain.Entity
	Name           string `json:"name"`
	Description    string `json:"description,omitempty"`
	Resource       string `json:"resource"` // e.g. "user", "chat", "document"
	Action         string `json:"action"`   // e.g. "create", "read", "update", "delete"
	Effect         string `json:"effect"`   // "allow" | "deny"
	OrganizationID string `json:"organization_id"`
}

func (p PolicyEntity) ToJSON() map[string]interface{} {
	// Marshal la struct a JSON
	data, err := json.Marshal(p)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}
	}

	// Unmarshal a map[string]interface{}
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}
	}

	return result
}
