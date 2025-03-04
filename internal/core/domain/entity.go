package domain

import (
	"encoding/json"
	"time"
)

// --------------------------------
// DOMAIN
// --------------------------------
// Entity
// --------------------------------

// Entity ahora implementa IEntity.
type Entity struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	IsRemoved bool      `json:"is_removed,omitempty"`
}

func (e Entity) GetID() string {
	return e.ID
}

func (r Entity) ToJSON() map[string]interface{} {
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
