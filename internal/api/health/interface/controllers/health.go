package controllers

import (
	"accounts/internal/common/logger"
	"accounts/internal/common/responses"
	"time"

	"github.com/gofiber/fiber/v2"
)

// HealthController estructura para manejar la ruta de Health
type HealthController struct {
}

// NewHealthController constructor para HealthController
func NewHealthController() *HealthController {
	return &HealthController{}
}

// GetHealth
func (c *HealthController) GetHealth(ctx *fiber.Ctx) error {

	entry := logger.GetByContext(ctx)

	entry.Info("HealthController.GetHealth")

	customResponse := responses.Response{
		Status: fiber.StatusOK,
		Data: fiber.Map{
			"status":    "ok",
			"message":   "El servicio está en línea y funcionando correctamente.",
			"timestamp": time.Now().Unix(),
		},
		Metadata: fiber.Map{
			"trace_id":  "d316a340-9c0a-419c-ad25-b7fefcdda3ce",
			"caller_id": "000000",
		},
		Errors: nil,
	}

	// Se almacena el objeto para que el middleware lo procese
	ctx.Locals("response", customResponse)
	return nil
}
