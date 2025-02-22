package memory

import (
	"accounts/internal/core/domain"
	"encoding/json"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// Model
// --------------------------------

// Model se restringe a tipos que cumplan con IEntity.
type Model[E domain.IEntity] struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;primary_key;"`
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
