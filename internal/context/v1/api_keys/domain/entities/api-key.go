package entities

import (
	"accounts/internal/core/domain"
	"time"
)

// --------------------------------
// DOMAIN
// --------------------------------
// APIKey Entity
// --------------------------------

// APIKey embebe a Entity, por lo que automáticamente implementa domain.IEntity.

type APIKeyEntity struct {
	domain.Entity

	// Identidad / metadata
	Name           string `json:"name"`
	Description    string `json:"description"`
	OrganizationID string `json:"organization_id"`
	KeyID          string `json:"key_id"` // no-secreto, para vinculación/operación
	Prefix         string `json:"prefix"` // prefijo de la SECRET (para búsquedas)

	// Llaves
	SecretKey      string `json:"secret_key,omitempty"`      // se muestra SOLO al crear
	PublishableKey string `json:"publishable_key,omitempty"` // se puede exponer al cliente
	SecretHash     string `json:"secret_hash,omitempty"`     // hash de la secret key

	// Estado
	IsActive    bool       `json:"is_active"`
	Environment string     `json:"environment"` // "live" | "test"
	Scopes      []string   `json:"scopes,omitempty"`
	LastUsedAt  *time.Time `json:"last_used_at,omitempty"`
	RevokedAt   *time.Time `json:"revoked_at,omitempty"`
}

func (o APIKeyEntity) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"id":              o.ID,
		"name":            o.Name,
		"description":     o.Description,
		"organization_id": o.OrganizationID,
		"key_id":          o.KeyID,
		"prefix":          o.Prefix,
		"secret_key":      o.SecretKey,
		"publishable_key": o.PublishableKey,
		"is_active":       o.IsActive,
		"environment":     o.Environment,
		"scopes":          o.Scopes,
		"last_used_at":    o.LastUsedAt,
		"revoked_at":      o.RevokedAt,
		"created_at":      o.CreatedAt,
		"updated_at":      o.UpdatedAt,
	}
}
