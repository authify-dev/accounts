package emails

import (
	"accounts/internal/api/v1/login_methods/domain/entities"
	"accounts/internal/db/postgres"
	postgres_users "accounts/internal/db/postgres/users"

	"github.com/google/uuid"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// User Model
// --------------------------------

// LoginMethodModel utiliza Model parametrizado con User.
type LoginMethodModel struct {
	postgres.Model[entities.LoginMethod]
	UserID     uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	EntityID   uuid.UUID `gorm:"type:uuid;not null" json:"entity_id"`
	EntityType string    `gorm:"type:varchar(255);not null" json:"entity_type"`
	IsActive   bool      `gorm:"type:boolean;not null" json:"is_active"`
	IsVerify   bool      `gorm:"type:boolean;not null" json:"is_verify"`

	// La etiqueta foreignKey indica cuál es el campo en este modelo que es llave foránea,
	// y references indica a qué campo del modelo relacionado hace referencia.
	UserModel postgres_users.UserModel `gorm:"foreignKey:UserID;references:ID" json:"user"`
}

func (LoginMethodModel) TableName() string {
	return "login_methods"
}

func (c LoginMethodModel) GetID() uuid.UUID {
	return c.ID
}
