package domain

import (
	"time"

	"github.com/google/uuid"
)

type Entity struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	IsRemoved bool      `json:"is_removed,omitempty"`
}
