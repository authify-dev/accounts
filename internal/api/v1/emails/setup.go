package emails

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"accounts/internal/api/v1/emails/domain/services"
	"accounts/internal/api/v1/emails/interface/controllers"
	utils_controller "accounts/internal/common/controllers"
	"accounts/internal/core/settings"
	codes "accounts/internal/db/postgres/codes"
	emails "accounts/internal/db/postgres/emails"
	login_methods "accounts/internal/db/postgres/login_methods"
	refresh "accounts/internal/db/postgres/refresh_tokens"
	roles "accounts/internal/db/postgres/role"
	users "accounts/internal/db/postgres/users"
)

func SetupEmailsModule(app *gin.Engine) {

	db, err := gorm.Open(postgres.Open(settings.Settings.POSTGRES_DSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	service := services.NewEmailsService(
		emails.NewEmailPostgresRepository(db),
		users.NewUserPostgresRepository(db),
		roles.NewRolePostgresRepository(db),
		login_methods.NewLoginMethodPostgresRepository(db),
		codes.NewCodePostgresRepository(db),
		refresh.NewRefreshTokenPostgresRepository(db),
		utils_controller.JWTController{
			PublicKey:  settings.Settings.PUBLIC_KEY_JWT,
			PrivateKey: settings.Settings.PRIVATE_KEY_JWT,
		},
		utils_controller.NewPasswordController(settings.Settings.SECRET_PASSWORD),
	)

	controller := controllers.NewEmailsController(*service)

	// Rutas de users
	group := app.Group("/api/v1/emails")

	group.POST("/signup", controller.SignUp)
	group.POST("/signup/resend-code", controller.SignUpResendCode)

	group.POST("/signin", controller.SignIn)
	group.POST("/signin/resend-code", controller.SignInResendCode)

	group.POST("/activate", controller.Activate)

	group.POST("/reset-password", controller.ResetPassword)
	group.POST("/reset-password/resend-code", controller.ResetPasswordResendCode)
}
