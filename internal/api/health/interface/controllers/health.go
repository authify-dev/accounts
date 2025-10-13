package controllers

import (
	"foundation/domain/customctx"
	"foundation/domain/logger"
	"foundation/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// HealthController estructura para manejar la ruta de Health
type HealthController struct {
}

// NewHealthController constructor para HealthController
func NewHealthController() *HealthController {
	return &HealthController{}
}

// GetHealth
func (c *HealthController) GetHealth(ctx *gin.Context) {

	entry := logger.FromContext(ctx.Request.Context())

	entry.Info("HealthController.GetHealth")

	cc := customctx.NewCustomContext(ctx.Request.Context())

	response := utils.Response[map[string]any]{
		StatusCode: http.StatusOK,
		Success:    true,
		Data: map[string]any{
			"status":    "ok",
			"message":   "El servicio está en línea y funcionando correctamente.",
			"timestamp": time.Now().Unix(),
		},
	}

	// Se almacena el objeto para que el middleware lo procese
	ctx.JSON(response.StatusCode, response.ToMapWithCustomContext(cc))
}
