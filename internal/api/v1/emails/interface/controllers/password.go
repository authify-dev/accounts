package controllers

import (
	"accounts/internal/common/responses"

	"github.com/gofiber/fiber/v2"
)

func (c *EmailsController) ResetPassword(ctx *fiber.Ctx) error {
	customResponse := responses.Response{
		Status: fiber.StatusOK,
		Data:   "ResetPassword",
	}

	// Se almacena el objeto para que el middleware lo procese
	ctx.Locals("response", customResponse)
	return nil
}

func (c *EmailsController) ResetPasswordResendCode(ctx *fiber.Ctx) error {
	customResponse := responses.Response{
		Status: fiber.StatusOK,
		Data:   "ResetPasswordResendCode",
	}

	// Se almacena el objeto para que el middleware lo procese
	ctx.Locals("response", customResponse)
	return nil
}
