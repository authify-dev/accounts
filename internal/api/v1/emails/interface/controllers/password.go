package controllers

import (
	"accounts/internal/common/responses"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func (c *EmailsController) ResetPassword(ctx *gin.Context) {
	customResponse := responses.Response{
		Status: fiber.StatusOK,
		Data:   "ResetPassword",
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
