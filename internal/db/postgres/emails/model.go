package postgres

import (
	"accounts/internal/api/v1/oauth_logins/domain/entities"
	"accounts/internal/db/postgres"
	postgres_users "accounts/internal/db/postgres/users"

	"github.com/google/uuid"
)

// EmailModel representa el modelo de datos para la entidad OAuthLogin.
type EmailModel struct {
	// Se asume que postgres.Model es un struct genérico que contiene campos comunes (como ID).
	postgres.Model[entities.OAuthLogin]

	// UserID es el identificador del usuario asociado.
	UserID uuid.UUID `gorm:"type:uuid;not null" json:"user_id,omitempty"`

	// ExternalID representa el identificador externo de la entidad.
	Email string `gorm:"type:varchar(255);uniqueIndex;not null" json:"email,omitempty"`

	// Platform indica la plataforma del login OAuth (por ejemplo, Google, Facebook, etc.).
	Password string `gorm:"type:varchar(255);not null" json:"password,omitempty"`

	// User es el usuario asociado al login OAuth.
	UserModel postgres_users.UserModel `gorm:"foreignKey:UserID;references:ID" json:"user"`
}

// TableName especifica el nombre de la tabla en la base de datos.
func (EmailModel) TableName() string {
	return "emails"
}

// GetID retorna el identificador único del modelo.
func (o EmailModel) GetID() uuid.UUID {
	return o.ID
}
