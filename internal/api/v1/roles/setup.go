package roles

import (
	"accounts/internal/api/v1/roles/domain/services"
	"accounts/internal/api/v1/roles/interface/controllers"
	"accounts/internal/core/settings"
	postgres_role "accounts/internal/db/postgres/role"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupRolesModule(app *fiber.App) {

	db, err := gorm.Open(postgres.Open(settings.Settings.POSTGRES_DSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println(db)

	rolesService := services.NewRolesService(
		postgres_role.NewRolePostgresRepository(db),
	)

	rolesController := controllers.NewRolesController(*rolesService)

	// Rutas de users
	roles := app.Group("/api/v1/roles")

	roles.Post("", rolesController.SignUp)
	roles.Get("", rolesController.List)

}
