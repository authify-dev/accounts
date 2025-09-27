package postgres

import (
	"accounts/internal/api/v1/codes/domain/entities"
	"accounts/internal/core/settings"
	postgres_users "accounts/internal/db/postgres/users"
	"foundation/infrastructure/db/cgorm"
	"foundation/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	ENTITY_CODE_CODE = "08"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// User Model
// --------------------------------

// CodeModel utiliza Model parametrizado con User.
type CodeModel struct {
	cgorm.Model[entities.Code]
	Code string `gorm:"type:varchar(255);not null;" json:"code"`

	UserID string `gorm:"type:varchar(50);not null" json:"user_id"`

	Type string `gorm:"type:varchar(50);not null" json:"type"`
	// La etiqueta foreignKey indica cuál es el campo en este modelo que es llave foránea,
	// y references indica a qué campo del modelo relacionado hace referencia.
	UserModel postgres_users.UserModel `gorm:"foreignKey:UserID;references:ID" json:"user"`
}

func (CodeModel) TableName() string {
	if settings.Settings.DB_SCHEMA == "" {
		return "codes"
	}
	return settings.Settings.DB_SCHEMA + ".codes"
}

func (c CodeModel) GetID() uuid.UUID {
	return c.ID
}

func (m *CodeModel) BeforeCreate(tx *gorm.DB) (err error) {
	idx, err := utils.NewUUIDx(settings.Settings.UUID_MODULE, ENTITY_CODE_CODE)
	if err != nil {
		return err
	}
	_id := idx.UUID()
	m.ID = _id
	return m.Model.BeforeCreate(tx)
}
