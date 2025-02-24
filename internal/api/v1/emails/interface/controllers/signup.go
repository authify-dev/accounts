package controllers

import (
	"accounts/internal/common/responses"

	"github.com/gofiber/fiber/v2"
)

func (c *EmailsController) SignUp(ctx *fiber.Ctx) error {
	customResponse := responses.Response{
		Status: fiber.StatusOK,
		Data:   "SignUp",
	}

	// Se almacena el objeto para que el middleware lo procese
	ctx.Locals("response", customResponse)
	return nil
}

func (c *EmailsController) SignUpResendCode(ctx *fiber.Ctx) error {
	customResponse := responses.Response{
		Status: fiber.StatusOK,
		Data:   "SignUpResendCode",
	}

	// Se almacena el objeto para que el middleware lo procese
	ctx.Locals("response", customResponse)
	return nil
}
