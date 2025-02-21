package middlewares

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CustomFormatter struct{}

// Format formatea la entrada del log con el formato "data | Level | trace-id | caller-id | log"
func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	traceID := "unknown"
	if v, ok := entry.Data["trace-id"]; ok {
		traceID = fmt.Sprintf("%v", v)
	}
	callerID := "unknown"
	if v, ok := entry.Data["caller-id"]; ok {
		callerID = fmt.Sprintf("%v", v)
	}
	logLine := fmt.Sprintf("%s | %s | %s | %s | %s\n",
		entry.Time.Format(time.RFC3339), // data: timestamp
		entry.Level.String(),            // Level
		traceID,                         // trace-id
		callerID,                        // caller-id
		entry.Message,                   // log
	)
	return []byte(logLine), nil
}

// LoggerMiddleware crea un logger contextualizado para cada petición usando logrus.
// Se espera que "trace-id" y "caller-id" ya estén almacenados en c.Locals.
func LoggerMiddleware(c *fiber.Ctx) error {
	// Recuperar "trace-id" y "caller-id" del contexto
	traceID, ok := c.Locals("trace-id").(string)
	if !ok || traceID == "" {
		traceID = "unknown"
	}
	callerID, ok := c.Locals("caller-id").(string)
	if !ok || callerID == "" {
		callerID = "unknown"
	}

	// Crear un nuevo logger y asignar el formateador personalizado
	logger := logrus.New()
	logger.SetFormatter(&CustomFormatter{})

	// Crear un entry con los campos contextuales
	entry := logger.WithFields(logrus.Fields{
		"trace-id":  traceID,
		"caller-id": callerID,
	})

	// Almacenar el logger en el contexto para que los handlers puedan usarlo
	c.Locals("logger", entry)

	entry.Infof("Request received: %s %s", c.Method(), c.OriginalURL())

	// Continuar con el siguiente middleware o handler
	return c.Next()
}
