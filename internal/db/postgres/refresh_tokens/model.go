package postgres

import (
	"accounts/internal/api/v1/refresh_tokens/domain/entities"
	"accounts/internal/db/postgres"
	postgres_login_methods "accounts/internal/db/postgres/login_methods"
	postgres_users "accounts/internal/db/postgres/users"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// User Model
// --------------------------------

// RefreshTokenModel utiliza Model parametrizado con User.
type RefreshTokenModel struct {
	postgres.Model[entities.RefreshToken]
	UserID        string `gorm:"type:varchar(50);not null" json:"user_id"`
	LoginMethodID string `gorm:"type:varchar(50);not null" json:"login_method_id,omitempty"`
	ExternalID    string `gorm:"type:varchar(50);not null" json:"external_id,omitempty"`

	ExpiresAt time.Time `json:"expires_at,omitempty"`
	RemoveAt  time.Time `json:"remove_at,omitempty"`

	// La etiqueta foreignKey indica cuál es el campo en este modelo que es llave foránea,
	// y references indica a qué campo del modelo relacionado hace referencia.
	UserModel        postgres_users.UserModel                `gorm:"foreignKey:UserID;references:ID" json:"user"`
	LoginMethodModel postgres_login_methods.LoginMethodModel `gorm:"foreignKey:LoginMethodID;references:ID" json:"login_method"`
}

func (RefreshTokenModel) TableName() string {
	return "refresh_tokens"
}

func (c RefreshTokenModel) GetID() string {
	return c.ID
}

func (m *RefreshTokenModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = fmt.Sprintf("%s_%s", m.TableName()[:3], uuid.New().String())
	return m.Model.BeforeCreate(tx)
}
