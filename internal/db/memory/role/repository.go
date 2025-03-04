package memory

import (
	"accounts/internal/api/v1/roles/domain/entities"
	"accounts/internal/core/domain"
	"accounts/internal/core/domain/criteria"
	"accounts/internal/db/memory"
	"fmt"
	"time"

	"github.com/google/uuid"
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
	id := fmt.Sprintf("%s_%s", "rol", uuid.New().String())
	r.roles = append(r.roles, RoleModel{
		Model: memory.Model[entities.Role]{
			ID:        id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			IsRemoved: false,
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

func (r *RoleMemoryRepository) Matching(criteria criteria.Criteria) ([]entities.Role, error) {

	return r.List()
}
