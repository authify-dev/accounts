package controllers

import (
	"accounts/internal/api/v1/emails/interface/dtos"
	"accounts/internal/common/logger"
	"accounts/internal/common/requests"
	"accounts/internal/common/responses"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func (c *EmailsController) ResetPassword(ctx *gin.Context) {

	entry := logger.FromContext(ctx.Request.Context())

	entry.Info("ResetPassword")

	dto := requests.GetDTO[dtos.ResetPasswordDTO](ctx)

	entity := dto.ToEntity()
	entry.Infof("DTO: %v", dto)
	entry.Infof("Entity: %v", entity)

	response := c.userService.ResetPassword(ctx.Request.Context(), entity)

	ctx.JSON(response.StatusCode, response.ToMap())
}

func (c *EmailsController) ResetPasswordConfirm(ctx *gin.Context) {
	customResponse := responses.Response{
		Status: fiber.StatusOK,
		Data:   "ResetPasswordConfirm",
	}

	// Se almacena el objeto para que el middleware lo procese
	ctx.JSON(fiber.StatusOK, customResponse)
}

func (c *EmailsController) ResetPasswordResendCode(ctx *gin.Context) {
	customResponse := responses.Response{
		Status: fiber.StatusOK,
		Data:   "ResetPasswordResendCode",
	}

	// Se almacena el objeto para que el middleware lo procese
	ctx.JSON(fiber.StatusOK, customResponse)
}
