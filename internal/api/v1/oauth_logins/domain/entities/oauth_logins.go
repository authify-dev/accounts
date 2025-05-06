package entities

import (
	"accounts/internal/core/domain"
)

// --------------------------------
// DOMAIN
// --------------------------------
// OAuthLogins Entity
// --------------------------------

// OAuthLogins embebe a Entity, por lo que automáticamente implementa domain.IEntity.
type OAuthLogin struct {
	domain.Entity
	UserID     string `json:"user_id,omitempty"`
	ExternalID string `json:"entity_id,omitempty"`
	Platform   string `json:"platform,omitempty"`
	Email      string `json:"email,omitempty"`
}
