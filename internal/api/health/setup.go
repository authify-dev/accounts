package health

import (
	"accounts/internal/api/health/interface/controllers"
	"accounts/internal/common/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupHealthModule(app *fiber.App) {

	healthController := controllers.NewHealthController()

	// Rutas de health
	health := app.Group("/health")

	health.Use(middlewares.CatcherMiddleware)

	health.Get("", healthController.GetHealth)
}
