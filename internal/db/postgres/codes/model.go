package postgres

import (
	"accounts/internal/api/v1/codes/domain/entities"
	"accounts/internal/db/postgres"
	postgres_users "accounts/internal/db/postgres/users"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// User Model
// --------------------------------

// CodeModel utiliza Model parametrizado con User.
type CodeModel struct {
	postgres.Model[entities.Code]
	Code string `gorm:"type:varchar(255);not null;" json:"code"`

	UserID string `gorm:"type:varchar(50);not null" json:"user_id"`
	// La etiqueta foreignKey indica cuál es el campo en este modelo que es llave foránea,
	// y references indica a qué campo del modelo relacionado hace referencia.
	UserModel postgres_users.UserModel `gorm:"foreignKey:UserID;references:ID" json:"user"`
}

func (CodeModel) TableName() string {
	return "codes"
}

func (c CodeModel) GetID() string {
	return c.ID
}

func (m *CodeModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = fmt.Sprintf("%s_%s", m.TableName()[:3], uuid.New().String())
	return m.Model.BeforeCreate(tx)
}
