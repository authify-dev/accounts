package postgres

import (
	"accounts/internal/api/v1/users/domain/entities"
	"accounts/internal/core/settings"
	postgres_role "accounts/internal/db/postgres/role"
	"foundation/infrastructure/db/cgorm"
	"foundation/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	ENTITY_USER_CODE = "05"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// User Model
// --------------------------------

// UserModel utiliza Model parametrizado con User.
type UserModel struct {
	cgorm.Model[entities.User]
	UserName string `gorm:"type:varchar(255);uniqueIndex;not null;" json:"user_name"`
	Name     string `gorm:"type:varchar(255);" json:"name"`

	RoleID string `gorm:"type:varchar(50);not null" json:"role_id"`
	// La etiqueta foreignKey indica cuál es el campo en este modelo que es llave foránea,
	// y references indica a qué campo del modelo relacionado hace referencia.
	RoleModel postgres_role.RoleModel `gorm:"foreignKey:RoleID;references:ID" json:"company"`
}

func (UserModel) TableName() string {
	if settings.Settings.DB_SCHEMA == "" {
		return "users"
	}
	return settings.Settings.DB_SCHEMA + ".users"
}

func (c UserModel) GetID() uuid.UUID {
	return c.ID
}

func (m *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	idx, err := utils.NewUUIDx(settings.Settings.UUID_MODULE, ENTITY_USER_CODE)
	if err != nil {
		return err
	}
	m.ID = idx.UUID()
	if err != nil {
		return err
	}
	return m.Model.BeforeCreate(tx)
}
