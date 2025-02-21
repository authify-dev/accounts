package middlewares

import (
	"accounts/internal/common/responses"

	"github.com/gofiber/fiber/v2"
)

// RealResponseMiddleware recupera el objeto CustomResponse y genera la respuesta final
func CatcherMiddleware(c *fiber.Ctx) error {
	// Tareas previas, por ejemplo, establecer cabeceras
	c.Set("X-Custom", "CustomMiddleware")

	// Ejecuta el siguiente handler (en este caso, GetHealth)
	if err := c.Next(); err != nil {
		return err
	}

	// Luego de ejecutar el handler, se recupera el objeto CustomResponse
	if resp, ok := c.Locals("response").(responses.Response); ok {
		return c.Status(resp.Status).JSON(resp)
	}
	// En caso de no encontrarse, se puede devolver un error o una respuesta por defecto
	return nil
}
