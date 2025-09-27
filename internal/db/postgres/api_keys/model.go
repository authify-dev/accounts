package apikeys_gorm

import (
	"accounts/internal/context/v1/api_keys/domain/entities"
	"accounts/internal/core/settings"
	organizations_gorm "accounts/internal/db/postgres/organinizations"
	"foundation/infrastructure/db/cgorm"
	"foundation/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	ENTITY_API_KEY_CODE = "0c"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// APIKey Model
// --------------------------------

// Tabla unificada: api_keys
type APIKeyModel struct {
	cgorm.Model[entities.APIKeyEntity]

	OrganizationID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:uniq_api_keys_org_name,priority:1;index:idx_api_keys_org_prefix,priority:1" json:"organization_id"`
	Name           string    `gorm:"type:varchar(255);not null;uniqueIndex:uniq_api_keys_org_name,priority:2" json:"name"`

	// Identificador y prefijo (para búsquedas por secret)
	KeyID  string `gorm:"type:varchar(32);not null;uniqueIndex" json:"key_id"`
	Prefix string `gorm:"type:varchar(16);not null;index:idx_api_keys_org_prefix,priority:2" json:"prefix"`

	// Secret: SÓLO hash en DB (Argon2id + pepper). Nunca guardes el plaintext.
	SecretHash string `gorm:"type:text;not null" json:"secret_hash"`

	// Publishable: puede guardarse en claro (o también hasheada si prefieres).
	PublishableKey string `gorm:"type:varchar(255);not null" json:"publishable_key"`

	// Metadatos
	Scopes      string     `gorm:"type:text" json:"scopes,omitempty"` // o JSONB
	Environment string     `gorm:"type:varchar(8);not null;default:live" json:"environment"`
	IsActive    bool       `gorm:"type:boolean;not null;default:true" json:"is_active"`
	LastUsedAt  *time.Time `json:"last_used_at,omitempty"`
	RevokedAt   *time.Time `json:"revoked_at,omitempty"`
	Description string     `gorm:"type:varchar(255)" json:"description,omitempty"`

	Organization *organizations_gorm.OrganizationModel `gorm:"foreignKey:OrganizationID;references:ID"`
}

func (APIKeyModel) TableName() string {
	if settings.Settings.DB_SCHEMA == "" {
		return "api_keys"
	}
	return settings.Settings.DB_SCHEMA + ".api_keys"
}

func (c APIKeyModel) GetID() uuid.UUID {
	return c.ID
}

func (m *APIKeyModel) BeforeCreate(tx *gorm.DB) (err error) {
	idx, err := utils.NewUUIDx(settings.Settings.UUID_MODULE, ENTITY_API_KEY_CODE)
	if err != nil {
		return err
	}
	m.ID = idx.UUID()
	return m.Model.BeforeCreate(tx)
}
