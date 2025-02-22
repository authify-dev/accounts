package memory

import (
	"accounts/internal/api/v1/roles/domain/entities"
	"accounts/internal/core/domain"
	"accounts/internal/db/memory"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// Role Memory Repository
// --------------------------------

type RoleMemoryRepository struct {
	memory.MemoryRepository[entities.Role]
	roles []RoleModel
}

func (r *RoleMemoryRepository) Save(role entities.Role) error {
	r.roles = append(r.roles, RoleModel{
		Model: memory.Model[entities.Role]{
			ID: uuid.New(),
			Model: gorm.Model{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
			},
		},
		Name:        role.Name,
		Description: role.Description,
	})

	return nil
}

func (r *RoleMemoryRepository) List() ([]entities.Role, error) {

	var rolesEntities []entities.Role

	for _, role := range r.roles {

		result := domain.ModelToEntity[entities.Role, RoleModel](role)

		if result.Err != nil {
			fmt.Println("Error al convertir a entidad:", result.Err)
			return nil, result.Err
		}

		rolesEntities = append(rolesEntities, result.Data)
	}

	return rolesEntities, nil
}
