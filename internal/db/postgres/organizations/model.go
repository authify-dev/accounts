package organizations_gorm

import (
	"accounts/internal/context/v1/organizations/domain/entities"
	"accounts/internal/core/settings"
	"accounts/internal/db/postgres"
	"foundation/types/uuidx"

	"gorm.io/gorm"
)

const (
	ENTITY_ORGANIZATION_CODE = "01"
)

// OrganizationModel representa el modelo de datos para la entidad Organization.
type OrganizationModel struct {
	// Se asume que postgres.Model es un struct genérico que contiene campos comunes (como ID).
	postgres.Model[entities.Organization]

	// RootUserID es el identificador del usuario raíz de la organización.
	RootUserID string `gorm:"type:varchar(50);not null" json:"root_user_id,omitempty"`
}

// TableName especifica el nombre de la tabla en la base de datos.
func (OrganizationModel) TableName() string {
	if settings.Settings.DB_SCHEMA == "" {
		return "organizations"
	}
	return settings.Settings.DB_SCHEMA + ".organizations"
}

// GetID retorna el identificador único del modelo.
func (o OrganizationModel) GetID() string {
	return o.ID
}

func (m *OrganizationModel) BeforeCreate(tx *gorm.DB) (err error) {
	idx, err := uuidx.NewUUIDx(settings.Settings.UUID_MODULE, ENTITY_ORGANIZATION_CODE)
	if err != nil {
		return err
	}
	m.ID = idx.UUID().String()

	return m.Model.BeforeCreate(tx)
}
