package controllers

import (
	"accounts/internal/common/responses"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func (c *RolesController) List(ctx *gin.Context) {

	organizationID := ctx.Query("organization_id")
	if organizationID == "" {
		ctx.JSON(fiber.StatusBadRequest, responses.Response{
			Status: fiber.StatusBadRequest,
			Errors: []string{"organization_id is required"},
		})
		return
	}
	roles, err := c.roles_service.List(organizationID)
	if err != nil {
		ctx.JSON(fiber.StatusBadRequest, responses.Response{
			Status: fiber.StatusBadRequest,
			Errors: []string{err.Error()},
		})
		return
	}

	customResponse := responses.Response{
		Status: fiber.StatusOK,
		Data:   roles,
	}

	// Se almacena el objeto para que el middleware lo procese
	ctx.JSON(fiber.StatusOK, customResponse)
}
