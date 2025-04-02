package entities

import (
	"accounts/internal/core/domain"
	"encoding/json"
)

// --------------------------------
// DOMAIN
// --------------------------------
// Codes Entity
// --------------------------------

// Codes embebe a Entity, por lo que automáticamente implementa domain.IEntity.
type Code struct {
	domain.Entity
	Code   string `json:"code,omitempty"`
	UserID string `json:"user_id,omitempty"`
	Type   string `json:"type,omitempty"`
	//User   string `json:"user,omitempty"`
}

func (r Code) ToJSON() map[string]interface{} {
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
