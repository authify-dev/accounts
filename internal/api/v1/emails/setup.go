package emails

import (
	"accounts/internal/api/v1/emails/interface/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupEmailsModule(app *fiber.App) {

	usersController := controllers.NewEmailsController()

	// Rutas de users
	users := app.Group("/api/v1/emails")

	users.Post("/signup", usersController.SignUp)
	users.Post("/signin", usersController.SignIn)
	users.Post("/activate", usersController.Activate)

}
