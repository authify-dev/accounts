package controllers

import (
	"accounts/internal/core/settings"

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

	// Responder con un JSON que contiene la URL generada
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"status":    "ok",
			"message":   "The service is online and functioning properly.",
			"timestamp": settings.Settings.TIMESTAMP,
		},
	})
}
