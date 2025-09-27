package middlewares

import (
	"context"
	"foundation/domain/logger"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware contextualiza el logger para una API REST en Gin.
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.GetString("trace-id")
		callerID := c.GetString("caller-id")

		fields := logger.LogFields{
			TraceID:  traceID,
			CallerID: callerID,
			Path:     c.Request.URL.Path,
			Method:   c.Request.Method,
			ClientIP: c.ClientIP(),
			UserID:   c.GetString("user_id"),
		}

		// Crear un logger contextualizado con información relevante.
		reqLogger := logger.WithFields(fields)

		// Loguear el inicio de la petición.
		reqLogger.Info("Inicio de petición API REST")

		// Agregar el logger al contexto de la request HTTP.
		ctx := context.WithValue(c.Request.Context(), "logger", reqLogger)
		ctx = context.WithValue(ctx, "fields", fields)
		c.Request = c.Request.WithContext(ctx)

		// Continuar con el siguiente middleware o handler.
		c.Next()

		// Loguear el fin de la petición.
		reqLogger.Info("Fin de petición API REST")
	}
}
