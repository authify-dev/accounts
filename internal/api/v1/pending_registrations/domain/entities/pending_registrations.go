package entities

import (
	"accounts/internal/core/domain"
	"encoding/json"
)

// --------------------------------
// DOMAIN
// --------------------------------
// Pending Registrations Entity
// --------------------------------

// PendingRegistrations embebe a Entity, por lo que automáticamente implementa domain.IEntity.
type PendingRegistration struct {
	domain.Entity
	UserName string `json:"user_name,omitempty"`
	Email    string `json:"email,omitempty"`
	Role     string `json:"role,omitempty"`
	Status   string `json:"status,omitempty"`
	CodeID   string `json:"code_id,omitempty"`
}

func (r PendingRegistration) ToJSON() map[string]interface{} {
	// Convertir el struct a JSON.
	data, err := json.Marshal(r)
	if err != nil {
		return nil
	}

	// Convertir los bytes JSON a un mapa.
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil
	}

	return result
}
