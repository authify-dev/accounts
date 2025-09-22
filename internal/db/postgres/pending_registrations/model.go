package postgres

import (
	"accounts/internal/api/v1/pending_registrations/domain/entities"
	"accounts/internal/core/settings"
	"accounts/internal/db/postgres"
	postgres_codes "accounts/internal/db/postgres/codes"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// PendingRegistrations Model
// --------------------------------

// PendingRegistrationModel utiliza Model parametrizado con User.
type PendingRegistrationModel struct {
	postgres.Model[entities.PendingRegistration]

	Email    string `gorm:"type:varchar(255);not null" json:"email"`
	UserName string `gorm:"type:varchar(255);not null" json:"user_name"`
	Role     string `gorm:"type:varchar(255);not null" json:"role"`
	Status   string `gorm:"type:varchar(255);not null" json:"status"`

	CodeID string `gorm:"type:uuid;not null" json:"code_id"`

	// La etiqueta foreignKey indica cuál es el campo en este modelo que es llave foránea,
	// y references indica a qué campo del modelo relacionado hace referencia.
	CodeModel postgres_codes.CodeModel `gorm:"foreignKey:CodeID;references:ID" json:"code"`
}

func (PendingRegistrationModel) TableName() string {
	if settings.Settings.DB_SCHEMA == "" {
		return "pending_registrations"
	}
	return settings.Settings.DB_SCHEMA + ".pending_registrations"
}

func (c PendingRegistrationModel) GetID() string {
	return c.ID
}

func (m *PendingRegistrationModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = fmt.Sprintf("%s_%s", m.TableName()[:3], uuid.New().String())
	return m.Model.BeforeCreate(tx)
}
