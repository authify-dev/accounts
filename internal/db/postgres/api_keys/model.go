package apikeys_gorm

import (
	"accounts/internal/context/v1/api_keys/domain/entities"
	"accounts/internal/core/settings"
	organizations_gorm "accounts/internal/db/postgres/organinizations"
	"foundation/infrastructure/db/cgorm"
	"foundation/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	ENTITY_API_KEY_CODE = "0c"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// Role Model
// --------------------------------

// RoleModel utiliza Model parametrizado con Role.
type APIKeyModel struct {
	cgorm.Model[entities.APIKeyEntity]

	// El name ya NO debe tener uniqueIndex “solo”
	Name string `gorm:"type:varchar(255);not null;uniqueIndex:idx_roles_org_name,priority:2" json:"name"`

	// Conviene que sea uuid si tu tabla de organizations usa uuid (ajústalo si aplica)
	OrganizationID string `gorm:"type:uuid;not null;uniqueIndex:idx_roles_org_name,priority:1" json:"organization_id"`

	Key string `gorm:"type:varchar(255);not null" json:"key"`

	Description string `gorm:"type:varchar(255);not null;" json:"description"`

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
