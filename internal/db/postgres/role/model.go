package postgres

import (
	"accounts/internal/api/v1/roles/domain/entities"
	"accounts/internal/db/postgres"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// Role Model
// --------------------------------

// RoleModel utiliza Model parametrizado con Role.
type RoleModel struct {
	postgres.Model[entities.Role]
	Name        string `gorm:"type:varchar(255);uniqueIndex;not null;" json:"name"`
	Description string `gorm:"type:varchar(255);not null;" json:"description"`
}

func (RoleModel) TableName() string {
	return "roles"
}

func (c RoleModel) GetID() string {
	return c.ID
}

func (m *RoleModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = fmt.Sprintf("%s_%s", m.TableName()[:3], uuid.New().String())
	return m.Model.BeforeCreate(tx)
}
