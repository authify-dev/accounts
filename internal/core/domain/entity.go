package domain

import (
	"time"

	"github.com/google/uuid"
)

// --------------------------------
// DOMAIN
// --------------------------------
// Entity
// --------------------------------

// Entity ahora implementa IEntity.
type Entity struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	IsRemoved bool      `json:"is_removed,omitempty"`
}

func (e Entity) GetID() uuid.UUID {
	return e.ID
}
