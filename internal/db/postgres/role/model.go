package postgres

import (
	"accounts/internal/api/v1/roles/domain/entities"
	"accounts/internal/core/settings"
	"accounts/internal/db/postgres"
	organizations_gorm "accounts/internal/db/postgres/organizations"
	"foundation/types/uuidx"

	"gorm.io/gorm"
)

const (
	ENTITY_ROLE_CODE = "03"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// Role Model
// --------------------------------

// RoleModel utiliza Model parametrizado con Role.
type RoleModel struct {
	postgres.Model[entities.Role]
	Name           string `gorm:"type:varchar(255);uniqueIndex;not null;" json:"name"`
	Description    string `gorm:"type:varchar(255);not null;" json:"description"`
	OrganizationID string `gorm:"type:varchar(50);not null" json:"organization_id"`

	Organization *organizations_gorm.OrganizationModel `gorm:"foreignKey:OrganizationID"`
}

func (RoleModel) TableName() string {
	if settings.Settings.DB_SCHEMA == "" {
		return "roles"
	}
	return settings.Settings.DB_SCHEMA + ".roles"
}

func (c RoleModel) GetID() string {
	return c.ID
}

func (m *RoleModel) BeforeCreate(tx *gorm.DB) (err error) {
	idx, err := uuidx.NewUUIDx(settings.Settings.UUID_MODULE, ENTITY_ROLE_CODE)
	if err != nil {
		return err
	}
	m.ID = idx.UUID().String()
	return m.Model.BeforeCreate(tx)
}
