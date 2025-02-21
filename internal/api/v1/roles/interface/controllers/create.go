package controllers

import (
	"accounts/internal/api/v1/roles/domain/entities"
	"accounts/internal/api/v1/roles/interface/dtos"
	"accounts/internal/common/requests"
	"accounts/internal/common/responses"

	"github.com/gofiber/fiber/v2"
)

func (c *RolesController) SignUp(ctx *fiber.Ctx) error {

	dto, err := requests.GetDTO[dtos.CreateRoleDTO](ctx)
	if err != nil {
		ctx.Locals("response", responses.Response{
			Status: fiber.StatusConflict,
			Errors: []string{err.Error()},
		})
		return nil
	}

	c.userService.Create(entities.Role{
		Name:        dto.Name,
		Description: dto.Description,
	})

	customResponse := responses.Response{
		Status: fiber.StatusOK,
		Data:   "Nuevo role",
	}

	// Se almacena el objeto para que el middleware lo procese
	ctx.Locals("response", customResponse)
	return nil
}
