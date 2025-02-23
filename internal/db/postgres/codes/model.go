package emails

import (
	"accounts/internal/api/v1/codes/domain/entities"
	"accounts/internal/db/postgres"
	postgres_users "accounts/internal/db/postgres/users"

	"github.com/google/uuid"
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

	UserID uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	// La etiqueta foreignKey indica cuál es el campo en este modelo que es llave foránea,
	// y references indica a qué campo del modelo relacionado hace referencia.
	UserModel postgres_users.UserModel `gorm:"foreignKey:UserID;references:ID" json:"user"`
}

func (CodeModel) TableName() string {
	return "codes"
}

func (c CodeModel) GetID() uuid.UUID {
	return c.ID
}
