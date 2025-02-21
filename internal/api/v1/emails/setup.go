package emails

import (
	"accounts/internal/api/v1/emails/interface/controllers"
	"accounts/internal/api/v1/users/domain/services"

	roles "accounts/internal/db/roles/postgres"
	users "accounts/internal/db/users/postgres"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupEmailsModule(app *fiber.App) {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	usersService := services.NewUsersService(
		&users.UserPostgresRepository{
			Conection: db,
		},
		&roles.RolePostgresRepository{
			Conection: db,
		},
	)

	usersController := controllers.NewEmailsController(*usersService)

	// Rutas de users
	users := app.Group("/api/v1/emails")

	users.Post("/signup", usersController.SignUp)
	users.Post("/signin", usersController.SignIn)
	users.Post("/activate", usersController.Activate)

}
