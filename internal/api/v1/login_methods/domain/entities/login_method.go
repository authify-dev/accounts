package entities

import (
	"accounts/internal/core/domain"
)

// --------------------------------
// DOMAIN
// --------------------------------
// LoginMethods Entity
// --------------------------------

// LoginMethods embebe a Entity, por lo que automáticamente implementa domain.IEntity.
type LoginMethod struct {
	domain.Entity
	UserID   string `json:"user_id,omitempty"`
	EntityID string `json:"entity_id,omitempty"`
}
