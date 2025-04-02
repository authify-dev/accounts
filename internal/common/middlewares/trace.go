package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Leer las cabeceras "trace-id" y "caller-id" de la petición
		traceID := c.GetHeader("trace-id")
		callerID := c.GetHeader("caller-id")

		// Si "trace-id" no existe, generar un UUID
		if traceID == "" {
			traceID = uuid.New().String()
		}

		// Si "caller-id" no existe, asignar "000000"
		if callerID == "" {
			callerID = "000000"
		}

		// Almacenar en el contexto para uso posterior en la cadena de handlers
		c.Set("trace-id", traceID)
		c.Set("caller-id", callerID)

		// Continuar con el siguiente middleware o handler
		c.Next()
	}
}
