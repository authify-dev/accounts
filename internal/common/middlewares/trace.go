package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func TraceMiddleware(c *fiber.Ctx) error {
	// Leer las cabeceras "trace-id" y "caller-id" de la petición
	traceID := c.Get("trace-id")
	callerID := c.Get("caller-id")

	// Si "trace-id" no existe, generar un UUID
	if traceID == "" {
		traceID = uuid.New().String()
	}

	// Si "caller-id" no existe, asignar "0000"
	if callerID == "" {
		callerID = "000000"
	}

	// Almacenar en el contexto usando c.Locals para uso posterior en la cadena de handlers
	c.Locals("trace-id", traceID)
	c.Locals("caller-id", callerID)

	// Continuar con el siguiente middleware o handler
	return c.Next()
}
