package postgres

import (
	"accounts/internal/api/v1/refresh_tokens/domain/entities"
	"accounts/internal/core/settings"
	postgres_login_methods "accounts/internal/db/postgres/login_methods"
	postgres_organizations "accounts/internal/db/postgres/organinizations"
	postgres_users "accounts/internal/db/postgres/users"
	"foundation/infrastructure/db/cgorm"
	"foundation/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	ENTITY_REFRESH_TOKEN_CODE = "0b"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// User Model
// --------------------------------

// RefreshTokenModel utiliza Model parametrizado con User.
type RefreshTokenModel struct {
	cgorm.Model[entities.RefreshToken]
	UserID        string `gorm:"type:varchar(50);not null" json:"user_id"`
	LoginMethodID string `gorm:"type:varchar(50);not null" json:"login_method_id,omitempty"`
	ExternalID    string `gorm:"type:varchar(50);not null" json:"external_id,omitempty"`

	ExpiresAt time.Time `json:"expires_at,omitempty"`
	RemoveAt  time.Time `json:"remove_at,omitempty"`

	OrganizationID string `gorm:"type:uuid;not null" json:"organization_id"`

	// La etiqueta foreignKey indica cuál es el campo en este modelo que es llave foránea,
	// y references indica a qué campo del modelo relacionado hace referencia.
	UserModel         postgres_users.UserModel                 `gorm:"foreignKey:UserID;references:ID" json:"user"`
	LoginMethodModel  postgres_login_methods.LoginMethodModel  `gorm:"foreignKey:LoginMethodID;references:ID" json:"login_method"`
	OrganizationModel postgres_organizations.OrganizationModel `gorm:"foreignKey:OrganizationID;references:ID" json:"organization"`
}

func (RefreshTokenModel) TableName() string {
	if settings.Settings.DB_SCHEMA == "" {
		return "refresh_tokens"
	}
	return settings.Settings.DB_SCHEMA + ".refresh_tokens"
}

func (c RefreshTokenModel) GetID() uuid.UUID {
	return c.ID
}

func (m *RefreshTokenModel) BeforeCreate(tx *gorm.DB) (err error) {
	idx, err := utils.NewUUIDx(settings.Settings.UUID_MODULE, ENTITY_REFRESH_TOKEN_CODE)
	if err != nil {
		return err
	}
	_id := idx.UUID()
	m.ID = _id
	return m.Model.BeforeCreate(tx)
}
