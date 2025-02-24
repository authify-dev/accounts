package middlewares

import (
	"accounts/internal/common/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// LoggerMiddleware contextualiza el logger para una API REST en Gin.
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generar un ID único para la petición.
		requestID := uuid.New().String()

		// Crear un logger contextualizado con información relevante.
		reqLogger := logger.WithFields(map[string]interface{}{
			"request_id": requestID,
			"session_id": "sessionID",
			"path":       c.Request.URL.Path,
			"method":     c.Request.Method,
		})

		// Loguear el inicio de la petición.
		reqLogger.Info("Inicio de petición API REST")

		// Agregar el logger contextualizado al contexto de Gin para uso posterior.
		c.Set("logger", reqLogger)

		// Continuar con el siguiente middleware o handler.
		c.Next()

		// Loguear el fin de la petición.
		reqLogger.Info("Fin de petición API REST")
	}
}
