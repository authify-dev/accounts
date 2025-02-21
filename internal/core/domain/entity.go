package domain

import (
	"time"

	"github.com/google/uuid"
)

type Entity struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsRemoved bool      `json:"is_removed"`
}
