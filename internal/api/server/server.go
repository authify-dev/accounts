package server

import (
	"accounts/internal/api/health"
	"accounts/internal/api/router"
	"accounts/internal/api/v1/emails"
	refreshtokens "accounts/internal/api/v1/refresh_tokens"
	"accounts/internal/api/v1/roles"
	"accounts/internal/api/v1/users"

	"accounts/internal/common/middlewares"
	"accounts/internal/core/settings"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

var Server *fiber.App

func Run() {

	app := setUpRouter()

	app.Run(fmt.Sprintf(":%d", settings.Settings.PORT))
}

func setUpRouter() *gin.Engine {

	app := router.NewRouter()

	app.Use(middlewares.TraceMiddleware())
	//app.Use(middlewares.CatcherMiddleware)
	app.Use(middlewares.LoggerMiddleware())

	health.SetupHealthModule(app)
	roles.SetupRolesModule(app)
	users.SetupUsersModule(app)
	emails.SetupEmailsModule(app)
	refreshtokens.SetupRefreshTokensModule(app)
	return app
}
