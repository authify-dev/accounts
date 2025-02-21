package server

import (
	"accounts/internal/api/health"
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

	health.SetupHealthModule(app)
	return app
}
