package main

import (
	"accounts/internal/api/v1/roles/domain/entities"
	"accounts/internal/api/v1/roles/domain/repositories"
	"accounts/internal/core/domain"
	"accounts/internal/db/memory"
	memory_role "accounts/internal/db/memory/role"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// --------------------------------
// DOMAIN
// --------------------------------
// UseCase
// --------------------------------

func UseCase(repo repositories.RoleRepository) {
	repo.Save(entities.Role{
		Name:        "Admin",
		Description: "Administrador",
		Entity: domain.Entity{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			IsRemoved: false,
		},
	})

	repo.Save(entities.Role{
		Name:        "User",
		Description: "Usuario",
		Entity: domain.Entity{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			IsRemoved: false,
		},
	})

	result := domain.ModelToEntity[entities.Role, memory_role.RoleModel](memory_role.RoleModel{
		Model: memory.Model[entities.Role]{
			ID: uuid.New(),
			Model: gorm.Model{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
			},
		},
		Name:        "Admin",
		Description: "Administrador",
	})

	if result.Err != nil {
		fmt.Println("Error al convertir a entidad:", result.Err)
		return
	}

	fmt.Println(result.Data)

	roles, err := repo.List()
	if err != nil {
		fmt.Println("Error al listar roles:", err)
		return
	}

	fmt.Println(len(roles))

	repo.View(roles)
}

// --------------------------------
// INTERFACE
// --------------------------------
// Controller
// --------------------------------
func main() {
	repo := &memory_role.RoleMemoryRepository{}
	UseCase(repo)

}
