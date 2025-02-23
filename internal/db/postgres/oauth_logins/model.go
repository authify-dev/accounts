package postgres

import (
	"accounts/internal/api/v1/oauth_logins/domain/entities"
	"accounts/internal/db/postgres"
	postgres_users "accounts/internal/db/postgres/users"

	"github.com/google/uuid"
)

// OAuthLoginModel representa el modelo de datos para la entidad OAuthLogin.
type OAuthLoginModel struct {
	// Se asume que postgres.Model es un struct genérico que contiene campos comunes (como ID).
	postgres.Model[entities.OAuthLogin]

	// UserID es el identificador del usuario asociado.
	UserID uuid.UUID `gorm:"type:varchar(255);not null" json:"user_id,omitempty"`

	// ExternalID representa el identificador externo de la entidad.
	ExternalID uuid.UUID `gorm:"type:varchar(255);uniqueIndex;not null" json:"entity_id,omitempty"`

	// IsActive indica si el login OAuth se encuentra activo.
	IsActive bool `gorm:"default:true" json:"is_active,omitempty"`

	// IsVerify indica si el login OAuth ha sido verificado.
	IsVerify bool `gorm:"default:false" json:"is_verify,omitempty"`

	// Platform indica la plataforma del login OAuth (por ejemplo, Google, Facebook, etc.).
	Platform string `gorm:"type:varchar(255);not null" json:"platform,omitempty"`

	UserModel postgres_users.UserModel `gorm:"foreignKey:UserID;references:ID" json:"user"`
}

// TableName especifica el nombre de la tabla en la base de datos.
func (OAuthLoginModel) TableName() string {
	return "oauth_logins"
}

// GetID retorna el identificador único del modelo.
func (o OAuthLoginModel) GetID() uuid.UUID {
	return o.ID
}
