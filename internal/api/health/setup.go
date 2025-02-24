package health

import (
	"accounts/internal/api/health/interface/controllers"

	"github.com/gin-gonic/gin"
)

func SetupHealthModule(app *gin.Engine) {

	healthController := controllers.NewHealthController()

	// Rutas de health
	health := app.Group("/health")

	health.GET("", healthController.GetHealth)
}
