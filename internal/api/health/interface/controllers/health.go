package controllers

import (
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

	customResponse := responses.Response{
		Status: fiber.StatusOK,
		Data: fiber.Map{
			"status":    "ok",
			"message":   "El servicio está en línea y funcionando correctamente.",
			"timestamp": time.Now().Unix(),
		},
		Metadata: nil,
		Errors:   nil,
	}

	// Se almacena el objeto para que el middleware lo procese
	ctx.Locals("response", customResponse)
	return nil
}
