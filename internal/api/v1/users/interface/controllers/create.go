package controllers

import (
	"accounts/internal/api/v1/users/domain/entities"
	"accounts/internal/api/v1/users/interface/dtos"
	"accounts/internal/common/requests"
	"accounts/internal/common/responses"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func (c *UsersController) Create(ctx *gin.Context) {

	dto := requests.GetDTO[dtos.CreateUserDTO](ctx)

	c.userService.Create(entities.User{
		Name:     dto.Name,
		UserName: dto.UserName,
		Role:     dto.Role,
	})

	customResponse := responses.Response{
		Status: fiber.StatusOK,
		Data:   "Nuevo role",
	}

	// Se almacena el objeto para que el middleware lo procese
	ctx.JSON(fiber.StatusOK, customResponse)
}
