package roles

import (
	"accounts/internal/api/v1/roles/domain/services"
	"accounts/internal/api/v1/roles/interface/controllers"
	role_repo "accounts/internal/db/roles/postgres"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupRolesModule(app *fiber.App) {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	rolesService := services.NewRolesService(
		&role_repo.RolePostgresRepository{
			Conection: db,
		},
	)

	usersController := controllers.NewRolesController(*rolesService)

	// Rutas de users
	users := app.Group("/api/v1/roles")

	users.Post("", usersController.SignUp)

}
