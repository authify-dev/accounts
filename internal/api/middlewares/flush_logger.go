package middlewares

import (
	"foundation/domain/logger"

	"github.com/gin-gonic/gin"
)

func FlushLogsOnFinishMiddleware(h *logger.LokiHook) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ejecuta handlers
		c.Next()
		// Al finalizar, forzar flush de lo pendiente
		h.Flush()
	}
}
