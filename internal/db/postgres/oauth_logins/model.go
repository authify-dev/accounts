package postgres

import (
	"accounts/internal/api/v1/oauth_logins/domain/entities"
	"accounts/internal/db/postgres"
	postgres_users "accounts/internal/db/postgres/users"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// OAuthLoginModel representa el modelo de datos para la entidad OAuthLogin.
type OAuthLoginModel struct {
	// Se asume que postgres.Model es un struct genérico que contiene campos comunes (como ID).
	postgres.Model[entities.OAuthLogin]

	// UserID es el identificador del usuario asociado.
	UserID string `gorm:"type:varchar(50);not null" json:"user_id,omitempty"`

	// ExternalID representa el identificador externo de la entidad.
	ExternalID string `gorm:"type:varchar(255);uniqueIndex;not null" json:"entity_id,omitempty"`

	// Platform indica la plataforma del login OAuth (por ejemplo, Google, Facebook, etc.).
	Platform string `gorm:"type:varchar(255);not null" json:"platform,omitempty"`

	UserModel postgres_users.UserModel `gorm:"foreignKey:UserID;references:ID" json:"user"`

	Email string `gorm:"type:varchar(255)" json:"email,omitempty"`
}

// TableName especifica el nombre de la tabla en la base de datos.
func (OAuthLoginModel) TableName() string {
	return "oauth_logins"
}

// GetID retorna el identificador único del modelo.
func (o OAuthLoginModel) GetID() string {
	return o.ID
}

func (m *OAuthLoginModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = fmt.Sprintf("%s_%s", m.TableName()[:3], uuid.New().String())
	return m.Model.BeforeCreate(tx)
}
