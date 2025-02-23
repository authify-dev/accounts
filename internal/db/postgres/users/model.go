package postgres

import (
	"accounts/internal/api/v1/users/domain/entities"
	"accounts/internal/db/postgres"
	postgres_role "accounts/internal/db/postgres/role"

	"github.com/google/uuid"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// User Model
// --------------------------------

// UserModel utiliza Model parametrizado con User.
type UserModel struct {
	postgres.Model[entities.User]
	UserName string `gorm:"type:varchar(255);uniqueIndex;not null;" json:"user_name"`
	Name     string `gorm:"type:varchar(255);" json:"name"`

	RoleID uuid.UUID `gorm:"type:uuid;not null" json:"role_id"`
	// La etiqueta foreignKey indica cuál es el campo en este modelo que es llave foránea,
	// y references indica a qué campo del modelo relacionado hace referencia.
	RoleModel postgres_role.RoleModel `gorm:"foreignKey:RoleID;references:ID" json:"company"`
}

func (UserModel) TableName() string {
	return "users"
}

func (c UserModel) GetID() uuid.UUID {
	return c.ID
}
