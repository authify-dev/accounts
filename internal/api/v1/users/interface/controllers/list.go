package controllers

import (
	"accounts/internal/common/responses"

	"github.com/gofiber/fiber/v2"
)

func (c *UsersController) List(ctx *fiber.Ctx) error {

	roles, err := c.userService.List()
	if err != nil {
		ctx.Locals("response", responses.Response{
			Status: fiber.StatusConflict,
			Errors: []string{err.Error()},
		})
		return nil
	}

	customResponse := responses.Response{
		Status: fiber.StatusOK,
		Data:   roles,
	}

	// Se almacena el objeto para que el middleware lo procese
	ctx.Locals("response", customResponse)
	return nil
}
