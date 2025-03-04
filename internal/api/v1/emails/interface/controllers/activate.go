package controllers

import (
	"accounts/internal/api/v1/emails/domain/entities"
	"accounts/internal/api/v1/emails/interface/dtos"
	"accounts/internal/common/requests"
	"accounts/internal/common/responses"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func (c *EmailsController) Activate(ctx *gin.Context) {
	dto := requests.GetDTO[dtos.ActivateDTO](ctx)

	entity, err := entities.NewActivateFromJSON(dto.ToJson())

	if err != nil {
		customResponse := responses.Response{
			Status: fiber.StatusBadRequest,
			Data:   "Error al parsear el JSON",
		}

		// Se almacena el objeto para que el middleware lo procese
		ctx.JSON(fiber.StatusOK, customResponse)
		return
	}

	response := c.userService.Activate(ctx.Request.Context(), entity)
	// Se almacena el objeto para que el middleware lo procese
	ctx.JSON(response.StatusCode, response.ToMap())
}
