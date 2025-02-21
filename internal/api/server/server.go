package server

import (
	"accounts/internal/api/health"
	"accounts/internal/api/v1/emails"

	"accounts/internal/common/middlewares"
	"accounts/internal/core/settings"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var Server *fiber.App

func Run() {

	app := setUpRouter()

	app.Listen(fmt.Sprintf(":%d", settings.Settings.PORT))
}

func setUpRouter() *fiber.App {

	app := fiber.New()

	app.Use(middlewares.CatcherMiddleware)
	app.Use(middlewares.TraceMiddleware)
	app.Use(middlewares.LoggerMiddleware)

	health.SetupHealthModule(app)
	emails.SetupEmailsModule(app)
	return app
}
