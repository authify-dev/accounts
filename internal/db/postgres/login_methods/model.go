package postgres

import (
	"accounts/internal/api/v1/login_methods/domain/entities"
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

// LoginMethodModel utiliza Model parametrizado con User.
type LoginMethodModel struct {
	postgres.Model[entities.LoginMethod]
	UserID     string `gorm:"type:varchar(50);not null" json:"user_id"`
	EntityID   string `gorm:"type:varchar(50);not null" json:"entity_id"`
	EntityType string `gorm:"type:varchar(255);not null" json:"entity_type"`
	IsActive   bool   `gorm:"type:boolean;not null" json:"is_active"`
	IsVerify   bool   `gorm:"type:boolean;not null" json:"is_verify"`

	// La etiqueta foreignKey indica cuál es el campo en este modelo que es llave foránea,
	// y references indica a qué campo del modelo relacionado hace referencia.
	UserModel postgres_users.UserModel `gorm:"foreignKey:UserID;references:ID" json:"user"`
}

func (LoginMethodModel) TableName() string {
	return "login_methods"
}

func (c LoginMethodModel) GetID() string {
	return c.ID
}

func (m *LoginMethodModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = fmt.Sprintf("%s_%s", m.TableName()[:3], uuid.New().String())
	return m.Model.BeforeCreate(tx)
}
