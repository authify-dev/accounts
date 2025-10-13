package entities

import (
	"accounts/internal/core/domain"
)

// --------------------------------
// DOMAIN
// --------------------------------
// Emails Entity
// --------------------------------

// Emails embebe a Entity, por lo que automáticamente implementa domain.IEntity.
type Email struct {
	domain.Entity
	UserID         string `json:"user_id,omitempty"`
	Email          string `json:"email,omitempty"`
	Password       string `json:"password,omitempty"`
	OrganizationID string `json:"organization_id,omitempty"`
}
