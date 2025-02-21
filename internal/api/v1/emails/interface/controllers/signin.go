package controllers

import (
	"accounts/internal/common/responses"

	"github.com/gofiber/fiber/v2"
)

func (c *EmailsController) SignIn(ctx *fiber.Ctx) error {

	customResponse := responses.Response{
		Status: fiber.StatusOK,
		Data:   "new inicio de sesion con email",
	}

	// Se almacena el objeto para que el middleware lo procese
	ctx.Locals("response", customResponse)
	return nil
}
