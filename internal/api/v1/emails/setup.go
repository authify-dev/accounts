package emails

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"accounts/internal/api/v1/emails/domain/services"
	"accounts/internal/api/v1/emails/interface/controllers"
	"accounts/internal/core/settings"
	emails "accounts/internal/db/postgres/emails"
	login_methods "accounts/internal/db/postgres/login_methods"
	users "accounts/internal/db/postgres/users"
)

func SetupEmailsModule(app *fiber.App) {

	db, err := gorm.Open(postgres.Open(settings.Settings.POSTGRES_DSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	service := services.NewEmailsService(
		emails.NewEmailPostgresRepository(db),
		users.NewUserPostgresRepository(db),
		login_methods.NewLoginMethodPostgresRepository(db),
	)

	controller := controllers.NewEmailsController(*service)

	// Rutas de users
	group := app.Group("/api/v1/emails")

	group.Post("/signup", controller.SignUp)
	group.Post("/signup/resend-code", controller.SignUpResendCode)

	group.Post("/signin", controller.SignIn)
	group.Post("/signin/resend-code", controller.SignInResendCode)

	group.Post("/activate", controller.Activate)

	group.Post("/reset-password", controller.ResetPassword)
	group.Post("/reset-password/resend-code", controller.ResetPasswordResendCode)
}
