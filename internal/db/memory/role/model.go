package memory

import (
	"accounts/internal/api/v1/roles/domain/entities"
	"accounts/internal/db/memory"

	"github.com/google/uuid"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// Role Model
// --------------------------------

// RoleModel utiliza Model parametrizado con Role.
type RoleModel struct {
	memory.Model[entities.Role]
	Name        string `gorm:"type:varchar(255);uniqueIndex;not null;" json:"name"`
	Description string `gorm:"type:varchar(255);not null;" json:"description"`
}

func (c RoleModel) GetID() uuid.UUID {
	return c.ID
}
