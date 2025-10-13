package postgres

import (
	"accounts/internal/api/v1/emails/domain/entities"
	"accounts/internal/core/settings"
	postgres_organizations "accounts/internal/db/postgres/organinizations"
	postgres_users "accounts/internal/db/postgres/users"
	"foundation/infrastructure/db/cgorm"
	"foundation/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	ENTITY_EMAIL_CODE = "06"
)

// EmailModel representa el modelo de datos para la entidad Email.
type EmailModel struct {
	// Se asume que postgres.Model es un struct genérico que contiene campos comunes (como ID).
	cgorm.Model[entities.Email]

	// UserID es el identificador del usuario asociado.
	UserID string `gorm:"type:varchar(50);not null" json:"user_id,omitempty"`

	// ExternalID representa el identificador externo de la entidad.
	Email string `gorm:"type:varchar(255);uniqueIndex:idx_emails_org_email,priority:2;not null" json:"email,omitempty"`

	// Platform indica la plataforma del login OAuth (por ejemplo, Google, Facebook, etc.).
	Password string `gorm:"type:varchar(255);not null" json:"password,omitempty"`

	// User es el usuario asociado al login OAuth.
	UserModel postgres_users.UserModel `gorm:"foreignKey:UserID;references:ID" json:"user"`

	OrganizationID    string                                   `gorm:"type:uuid;not null;uniqueIndex:idx_emails_org_email,priority:1" json:"organization_id"`
	OrganizationModel postgres_organizations.OrganizationModel `gorm:"foreignKey:OrganizationID;references:ID" json:"organization"`
}

// TableName especifica el nombre de la tabla en la base de datos.
func (EmailModel) TableName() string {
	if settings.Settings.DB_SCHEMA == "" {
		return "emails"
	}
	return settings.Settings.DB_SCHEMA + ".emails"
}

// GetID retorna el identificador único del modelo.
func (o EmailModel) GetID() uuid.UUID {
	return o.ID
}

func (m *EmailModel) BeforeCreate(tx *gorm.DB) (err error) {
	idx, err := utils.NewUUIDx(settings.Settings.UUID_MODULE, ENTITY_EMAIL_CODE)
	if err != nil {
		return err
	}
	m.ID = idx.UUID()
	if err != nil {
		return err
	}
	return m.Model.BeforeCreate(tx)
}
