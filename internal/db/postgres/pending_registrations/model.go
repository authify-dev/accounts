package postgres

import (
	"accounts/internal/api/v1/pending_registrations/domain/entities"
	"accounts/internal/core/settings"
	postgres_codes "accounts/internal/db/postgres/codes"
	postgres_organizations "accounts/internal/db/postgres/organinizations"
	"foundation/infrastructure/db/cgorm"
	"foundation/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	ENTITY_PENDING_REGISTRATION_CODE = "0a"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// PendingRegistrations Model
// --------------------------------

// PendingRegistrationModel utiliza Model parametrizado con User.
type PendingRegistrationModel struct {
	cgorm.Model[entities.PendingRegistration]

	Email    string `gorm:"type:varchar(255);not null" json:"email"`
	UserName string `gorm:"type:varchar(255);not null" json:"user_name"`
	Role     string `gorm:"type:varchar(255);not null" json:"role"`
	Status   string `gorm:"type:varchar(255);not null" json:"status"`

	CodeID         string `gorm:"type:uuid;not null" json:"code_id"`
	OrganizationID string `gorm:"type:uuid;not null" json:"organization_id"`

	// La etiqueta foreignKey indica cuál es el campo en este modelo que es llave foránea,
	// y references indica a qué campo del modelo relacionado hace referencia.
	CodeModel         postgres_codes.CodeModel                 `gorm:"foreignKey:CodeID;references:ID" json:"code"`
	OrganizationModel postgres_organizations.OrganizationModel `gorm:"foreignKey:OrganizationID;references:ID" json:"organization"`
}

func (PendingRegistrationModel) TableName() string {
	if settings.Settings.DB_SCHEMA == "" {
		return "pending_registrations"
	}
	return settings.Settings.DB_SCHEMA + ".pending_registrations"
}

func (c PendingRegistrationModel) GetID() uuid.UUID {
	return c.ID
}

func (m *PendingRegistrationModel) BeforeCreate(tx *gorm.DB) (err error) {
	idx, err := utils.NewUUIDx(settings.Settings.UUID_MODULE, ENTITY_PENDING_REGISTRATION_CODE)
	if err != nil {
		return err
	}
	_id := idx.UUID()
	m.ID = _id
	return m.Model.BeforeCreate(tx)
}
