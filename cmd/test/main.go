package main

import (
	role_entities "accounts/internal/api/v1/roles/domain/entities"
	"accounts/internal/api/v1/roles/domain/repositories"
	organization_entities "accounts/internal/context/v1/organizations/domain/entities"
	"accounts/internal/core/settings"
	postgres_organizations "accounts/internal/db/postgres/organinizations"
	postgres_role "accounts/internal/db/postgres/role"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// --------------------------------
// DOMAIN
// --------------------------------
// UseCase
// --------------------------------

func UseCase(repo repositories.RoleRepository, organizationRepo *postgres_organizations.OrganizationsPostgresRepository) {

	res := organizationRepo.Save(organization_entities.Organization{
		Name: "Organization 1",
	})

	if res.Err != nil {
		fmt.Println(res.Err)
	}

	fmt.Println(res.Data)

	resRole := repo.Save(role_entities.Role{
		Name:           "Admin",
		Description:    "Administrador",
		OrganizationID: res.Data.ID.String(),
	})

	if resRole.Err != nil {
		fmt.Println(resRole.Err)
	}

	fmt.Println(resRole.Data)

	// repo.Save(entities.Role{
	// 	Name:        "User",
	// 	Description: "Usuario",
	// 	Entity: domain.Entity{
	// 		ID:        uuid.New(),
	// 		CreatedAt: time.Now(),
	// 		UpdatedAt: time.Now(),
	// 		IsRemoved: false,
	// 	},
	// })

	// cri := criteria.Criteria{
	// 	Filters: *criteria.NewFilters(
	// 		[]criteria.Filter{
	// 			{
	// 				Field:    "name",
	// 				Operator: criteria.OperatorEqual,
	// 				Value:    "user",
	// 			},
	// 		},
	// 	),
	// }

	// repo.Matching(cri)
}

// --------------------------------
// INTERFACE
// --------------------------------
// Controller
// --------------------------------
func main() {
	settings.LoadDotEnv()
	settings.LoadEnvs()
	// Define el DSN para la conexión a PostgreSQL
	dsn := settings.Settings.POSTGRES_DSN

	// Conecta a la base de datos Postgres
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	repo := postgres_role.NewRolePostgresRepository(db)
	organizationRepo := postgres_organizations.NewOrganizationsPostgresRepository(db)
	UseCase(repo, organizationRepo)

}
