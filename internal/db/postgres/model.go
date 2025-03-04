package postgres

import (
	"accounts/internal/core/domain"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// Model
// --------------------------------

// Model se restringe a tipos que cumplan con IEntity.
type Model[E domain.IEntity] struct {
	ID        string    `gorm:"type:varchar(50);primary_key;" json:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	IsRemoved bool      `gorm:"type:boolean;default:false" json:"is_removed,omitempty"`
}

// BeforeCreate se ejecuta antes de insertar el registro y establece CreatedAt y UpdatedAt en UTC.
func (m *Model[E]) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().UTC()
	m.CreatedAt = now
	m.UpdatedAt = now
	return nil
}

// BeforeUpdate se ejecuta antes de actualizar el registro y establece UpdatedAt en UTC.
func (m *Model[E]) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now().UTC()
	return nil
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

func (c Model[E]) GetID() string {
	return c.ID
}

func (c Model[E]) TableName() string {
	return "model"
}
