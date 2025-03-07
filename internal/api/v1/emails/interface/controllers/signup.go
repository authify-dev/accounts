package controllers

import (
	"accounts/internal/api/v1/emails/domain/entities"
	"accounts/internal/api/v1/emails/interface/dtos"
	"accounts/internal/common/requests"
	"accounts/internal/common/responses"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func (c *EmailsController) SignUp(ctx *gin.Context) {

	dto := requests.GetDTO[dtos.SignUpDTO](ctx)

	entity, err := entities.NewSingUpFromJSON(dto.ToJson())

	if err != nil {
		customResponse := responses.Response{
			Status: fiber.StatusBadRequest,
			Data:   "Error al parsear el JSON",
		}

		// Se almacena el objeto para que el middleware lo procese
		ctx.JSON(fiber.StatusOK, customResponse)
		return
	}

	response := c.userService.SignUp(ctx.Request.Context(), entity)
	// Se almacena el objeto para que el middleware lo procese
	ctx.JSON(response.StatusCode, response.ToMap())
}

func (c *EmailsController) SignUpResendCode(ctx *gin.Context) {
	dto := requests.GetDTO[dtos.ResendActivationCodeDTO](ctx)

	entity, err := entities.NewResendActivationCodeFromJSON(dto.ToJson())

	if err != nil {
		customResponse := responses.Response{
			Status: fiber.StatusBadRequest,
			Data:   "Error al parsear el JSON",
		}

		// Se almacena el objeto para que el middleware lo procese
		ctx.JSON(fiber.StatusOK, customResponse)
		return
	}

	response := c.userService.ResendActivationCode(ctx.Request.Context(), entity)
	// Se almacena el objeto para que el middleware lo procese
	ctx.JSON(response.StatusCode, response.ToMap())
}
