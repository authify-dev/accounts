package organizations_gorm

import (
	"accounts/internal/context/v1/organizations/domain/entities"
	"accounts/internal/core/settings"
	"foundation/infrastructure/db/cgorm"
	"foundation/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	ENTITY_ORGANIZATION_CODE = "01"
)

// OrganizationModel representa el modelo de datos para la entidad Organization.
type OrganizationModel struct {
	cgorm.Model[entities.Organization]

	Name string `gorm:"type:varchar(255);not null;uniqueIndex:uniq_root_user_name,priority:2" json:"name"`

	RootUserID string `gorm:"type:varchar(50);not null;uniqueIndex:uniq_root_user_name,priority:1" json:"root_user_id"`
}

// TableName especifica el nombre de la tabla en la base de datos.
func (OrganizationModel) TableName() string {
	if settings.Settings.DB_SCHEMA == "" {
		return "organizations"
	}
	return settings.Settings.DB_SCHEMA + ".organizations"
}

// GetID retorna el identificador único del modelo.
func (o OrganizationModel) GetID() uuid.UUID {
	return o.ID
}

func (m *OrganizationModel) BeforeCreate(tx *gorm.DB) (err error) {
	idx, err := utils.NewUUIDx(settings.Settings.UUID_MODULE, ENTITY_ORGANIZATION_CODE)
	if err != nil {
		return err
	}
	m.ID = idx.UUID()

	return m.Model.BeforeCreate(tx)
}
