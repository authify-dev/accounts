package main

import (
	"accounts/internal/api/v1/roles/domain/entities"
	"accounts/internal/api/v1/roles/domain/repositories"
	"accounts/internal/core/domain"
	"accounts/internal/core/domain/criteria"
	postgres_role "accounts/internal/db/postgres/role"
	"time"

	"gorm.io/driver/sqlite"
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
			ID:        "uuid.New()",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			IsRemoved: false,
		},
	})

	repo.Save(entities.Role{
		Name:        "User",
		Description: "Usuario",
		Entity: domain.Entity{
			ID:        "uuid.New()",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			IsRemoved: false,
		},
	})

	cri := criteria.Criteria{
		Filters: *criteria.NewFilters(
			[]criteria.Filter{
				{
					Field:    "name",
					Operator: criteria.OperatorEqual,
					Value:    "user",
				},
			},
		),
	}

	repo.Matching(cri)
}

// --------------------------------
// INTERFACE
// --------------------------------
// Controller
// --------------------------------
func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	repo := postgres_role.NewRolePostgresRepository(db)
	UseCase(repo)

}
