package controllers

import (
	"accounts/internal/api/v1/emails/interface/dtos"
	"accounts/internal/api/v1/users/domain/entities"
	"accounts/internal/common/requests"
	"accounts/internal/common/responses"

	"github.com/gofiber/fiber/v2"
)

func (c *EmailsController) SignUp(ctx *fiber.Ctx) error {

	dto, err := requests.GetDTO[dtos.SignUpDTO](ctx)
	if err != nil {
		ctx.Locals("response", responses.Response{
			Status: fiber.StatusConflict,
			Errors: []string{err.Error()},
		})
		return nil
	}

	c.userService.Create(entities.User{
		Name:     dto.Name,
		UserName: dto.UserName,
		RoleID:   dto.Role,
	})

	customResponse := responses.Response{
		Status: fiber.StatusOK,
		Data:   "new registro de usuario con email",
	}

	// Se almacena el objeto para que el middleware lo procese
	ctx.Locals("response", customResponse)
	return nil
}
