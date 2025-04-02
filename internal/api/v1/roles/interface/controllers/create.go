package controllers

import (
	"accounts/internal/api/v1/roles/domain/entities"
	"accounts/internal/api/v1/roles/interface/dtos"
	"accounts/internal/common/requests"
	"accounts/internal/common/responses"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func (c *RolesController) SignUp(ctx *gin.Context) {

	dto := requests.GetDTO[dtos.CreateRoleDTO](ctx)

	c.userService.Create(entities.Role{
		Name:        dto.Name,
		Description: dto.Description,
	})

	customResponse := responses.Response{
		Status: fiber.StatusOK,
		Data:   "Nuevo role",
	}

	// Se almacena el objeto para que el middleware lo procese
	ctx.JSON(fiber.StatusOK, customResponse)
}
