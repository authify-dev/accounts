package entities

import (
	"accounts/internal/core/domain"
	"time"
)

// --------------------------------
// DOMAIN
// --------------------------------
// Refreshs Entity
// --------------------------------

// Refreshs embebe a Entity, por lo que automáticamente implementa domain.IEntity.
type RefreshToken struct {
	domain.Entity
	UserID        string    `json:"user_id,omitempty"`
	LoginMethodID string    `json:"login_method_id,omitempty"`
	ExternalID    string    `json:"external_id,omitempty"`
	ExpiresAt     time.Time `json:"expires_at,omitempty"`
	RemoveAt      time.Time `json:"remove_at,omitempty"`
}
