package postgres

import (
	"accounts/internal/api/v1/login_methods/domain/entities"
	"accounts/internal/core/settings"
	postgres_organizations "accounts/internal/db/postgres/organinizations"
	postgres_users "accounts/internal/db/postgres/users"
	"foundation/infrastructure/db/cgorm"
	"foundation/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	ENTITY_LOGIN_METHOD_CODE = "07"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// User Model
// --------------------------------

// LoginMethodModel utiliza Model parametrizado con User.
type LoginMethodModel struct {
	cgorm.Model[entities.LoginMethod]
	UserID     string `gorm:"type:varchar(50);not null" json:"user_id"`
	EntityID   string `gorm:"type:varchar(50);not null" json:"entity_id"`
	EntityType string `gorm:"type:varchar(255);not null" json:"entity_type"`
	IsActive   bool   `gorm:"type:boolean;not null" json:"is_active"`
	IsVerify   bool   `gorm:"type:boolean;not null" json:"is_verify"`

	OrganizationID string `gorm:"type:uuid;not null" json:"organization_id"`

	// La etiqueta foreignKey indica cuál es el campo en este modelo que es llave foránea,
	// y references indica a qué campo del modelo relacionado hace referencia.
	UserModel         postgres_users.UserModel                 `gorm:"foreignKey:UserID;references:ID" json:"user"`
	OrganizationModel postgres_organizations.OrganizationModel `gorm:"foreignKey:OrganizationID;references:ID" json:"organization"`
}

func (LoginMethodModel) TableName() string {
	if settings.Settings.DB_SCHEMA == "" {
		return "login_methods"
	}
	return settings.Settings.DB_SCHEMA + ".login_methods"
}

func (c LoginMethodModel) GetID() uuid.UUID {
	return c.ID
}

func (m *LoginMethodModel) BeforeCreate(tx *gorm.DB) (err error) {
	idx, err := utils.NewUUIDx(settings.Settings.UUID_MODULE, ENTITY_LOGIN_METHOD_CODE)
	if err != nil {
		return err
	}
	_id := idx.UUID()
	m.ID = _id
	return m.Model.BeforeCreate(tx)
}
