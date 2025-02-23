package postgres

import (
	"accounts/internal/core/domain"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// Model
// --------------------------------

// Model se restringe a tipos que cumplan con IEntity.
type Model[E domain.IEntity] struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	IsRemoved bool      `gorm:"type:boolean" json:"is_removed,omitempty"`
}

func (c *Model[E]) ToJSON() map[string]interface{} {
	var result map[string]interface{}

	// Convertir el struct a JSON (bytes).
	data, err := json.Marshal(c)
	if err != nil {
		// Manejo de error: se puede retornar un mapa vacío o nil.
		return nil
	}

	// Convertir los bytes JSON a un mapa.
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil
	}

	return result
}

func (c Model[E]) GetID() uuid.UUID {
	return c.ID
}
