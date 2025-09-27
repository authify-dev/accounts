package postgres

import (
	"accounts/internal/api/v1/oauth_logins/domain/entities"
	"accounts/internal/core/settings"
	postgres_organizations "accounts/internal/db/postgres/organinizations"
	postgres_users "accounts/internal/db/postgres/users"
	"foundation/infrastructure/db/cgorm"
	"foundation/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	ENTITY_OAUTH_LOGIN_CODE = "09"
)

// OAuthLoginModel representa el modelo de datos para la entidad OAuthLogin.
type OAuthLoginModel struct {
	// Se asume que postgres.Model es un struct genérico que contiene campos comunes (como ID).
	cgorm.Model[entities.OAuthLogin]

	// UserID es el identificador del usuario asociado.
	UserID string `gorm:"type:varchar(50);not null" json:"user_id,omitempty"`

	// ExternalID representa el identificador externo de la entidad.
	ExternalID string `gorm:"type:varchar(255);uniqueIndex;not null" json:"entity_id,omitempty"`

	// Platform indica la plataforma del login OAuth (por ejemplo, Google, Facebook, etc.).
	Platform string `gorm:"type:varchar(255);not null" json:"platform,omitempty"`

	UserModel         postgres_users.UserModel                 `gorm:"foreignKey:UserID;references:ID" json:"user"`
	OrganizationModel postgres_organizations.OrganizationModel `gorm:"foreignKey:OrganizationID;references:ID" json:"organization"`

	OrganizationID string `gorm:"type:uuid;not null" json:"organization_id"`

	Email string `gorm:"type:varchar(255)" json:"email,omitempty"`
}

// TableName especifica el nombre de la tabla en la base de datos.
func (OAuthLoginModel) TableName() string {
	if settings.Settings.DB_SCHEMA == "" {
		return "oauth_logins"
	}
	return settings.Settings.DB_SCHEMA + ".oauth_logins"
}

// GetID retorna el identificador único del modelo.
func (o OAuthLoginModel) GetID() uuid.UUID {
	return o.ID
}

func (m *OAuthLoginModel) BeforeCreate(tx *gorm.DB) (err error) {
	idx, err := utils.NewUUIDx(settings.Settings.UUID_MODULE, ENTITY_OAUTH_LOGIN_CODE)
	if err != nil {
		return err
	}
	_id := idx.UUID()
	m.ID = _id
	return m.Model.BeforeCreate(tx)
}
