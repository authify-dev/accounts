package controllers

import (
	"accounts/internal/api/v1/users/domain/entities"
	"accounts/internal/common/responses"

	"github.com/gofiber/fiber/v2"
)

func (c *EmailsController) SignUp(ctx *fiber.Ctx) error {

	c.userService.Create(entities.User{
		Name: "nombre",
	})

	customResponse := responses.Response{
		Status: fiber.StatusOK,
		Data:   "new registro de usuario con email",
	}

	// Se almacena el objeto para que el middleware lo procese
	ctx.Locals("response", customResponse)
	return nil
}
