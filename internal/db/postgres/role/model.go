package postgres

import (
	"accounts/internal/api/v1/roles/domain/entities"
	"accounts/internal/core/settings"
	organizations_gorm "accounts/internal/db/postgres/organinizations"
	"foundation/infrastructure/db/cgorm"
	"foundation/utils"

	"github.com/google/uuid"
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
	cgorm.Model[entities.Role]

	// El name ya NO debe tener uniqueIndex “solo”
	Name string `gorm:"type:varchar(255);not null;uniqueIndex:idx_roles_org_name,priority:2" json:"name"`

	// Conviene que sea uuid si tu tabla de organizations usa uuid (ajústalo si aplica)
	OrganizationID string `gorm:"type:uuid;not null;uniqueIndex:idx_roles_org_name,priority:1" json:"organization_id"`

	Description string `gorm:"type:varchar(255);not null;" json:"description"`

	Organization *organizations_gorm.OrganizationModel `gorm:"foreignKey:OrganizationID;references:ID"`
}

func (RoleModel) TableName() string {
	if settings.Settings.DB_SCHEMA == "" {
		return "roles"
	}
	return settings.Settings.DB_SCHEMA + ".roles"
}

func (c RoleModel) GetID() uuid.UUID {
	return c.ID
}

func (m *RoleModel) BeforeCreate(tx *gorm.DB) (err error) {
	idx, err := utils.NewUUIDx(settings.Settings.UUID_MODULE, ENTITY_ROLE_CODE)
	if err != nil {
		return err
	}
	m.ID = idx.UUID()
	return m.Model.BeforeCreate(tx)
}
