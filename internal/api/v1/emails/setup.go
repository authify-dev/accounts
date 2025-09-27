package emails

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"accounts/internal/api/v1/emails/domain/services"
	"accounts/internal/api/v1/emails/interface/controllers"
	utils_controller "accounts/internal/common/controllers"
	"accounts/internal/core/domain/event"
	"accounts/internal/core/infrastructure/event_bus/local"
	"accounts/internal/core/infrastructure/event_bus/rabbit"
	"accounts/internal/core/settings"
	codes "accounts/internal/db/postgres/codes"
	emails "accounts/internal/db/postgres/emails"
	login_methods "accounts/internal/db/postgres/login_methods"
	pending_registrations "accounts/internal/db/postgres/pending_registrations"
	refresh "accounts/internal/db/postgres/refresh_tokens"
	roles "accounts/internal/db/postgres/role"
	users "accounts/internal/db/postgres/users"
)

func SetupEmailsModule(app *gin.Engine, db *gorm.DB) {

	service := services.NewEmailsService(
		emails.NewEmailPostgresRepository(db),
		pending_registrations.NewPendingRegistrationsPostgresRepository(db),
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
		LocalEventBus(),
	)

	controller := controllers.NewEmailsController(*service)

	// Rutas de users
	group := app.Group(settings.Settings.ROOT_PATH + "/:organization_id/api/v1/emails")

	group.POST("/signup", controller.SignUp)
	group.POST("/signup/resend-code", controller.SignUpResendCode)

	group.POST("/signin", controller.SignIn)
	group.POST("/signin/resend-code", controller.SignInResendCode)

	group.POST("/activate", controller.Activate)
	group.POST("/activate/v2", controller.ActivateV2)
	group.POST("/activate/v2/set-password", controller.SetPassword)

	group.POST("/reset", controller.ResetPassword)
	group.POST("/reset-confirm", controller.ResetPasswordConfirm)
}

func eventBus() event.EventBus {

	connection := rabbit.NewRabbitMqConnection(
		event.SettingsEventBus{
			Username: settings.Settings.USER_EVENT_BUS,
			Password: settings.Settings.PASSWORD_EVENT_BUS,
			VHost:    settings.Settings.VHOST_EVENT_BUS,
			Connection: struct {
				Hostname string
				Port     int
			}{
				Hostname: settings.Settings.HOST_EVENT_BUS,
				Port:     settings.Settings.PORT_EVENT_BUS,
			},
		},
	)

	connection.Connect()

	event_bus := rabbit.NewRabbitEventBus(
		connection,
		"domain_events",
	)

	return event_bus
}

func LocalEventBus() event.EventBus {
	return local.MockEventBus()
}
