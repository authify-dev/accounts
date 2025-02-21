package logger

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func GetByContext(ctx *fiber.Ctx) *logrus.Entry {
	if logger, ok := ctx.Locals("logger").(*logrus.Entry); ok {
		return logger
	}
	return nil
}
