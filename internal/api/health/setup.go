package health

import (
	"accounts/internal/api/health/interface/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupHealthModule(app *fiber.App) {

	healthController := controllers.NewHealthController()

	// Rutas de health
	health := app.Group("/health")

	health.Get("", healthController.GetHealth)
}
