package controllers

import (
	"accounts/internal/common/responses"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func (c *EmailsController) SignIn(ctx *gin.Context) {
	customResponse := responses.Response{
		Status: fiber.StatusOK,
		Data:   "SignIn",
	}

	// Se almacena el objeto para que el middleware lo procese
	ctx.JSON(fiber.StatusOK, customResponse)
}

func (c *EmailsController) SignInResendCode(ctx *gin.Context) {
	customResponse := responses.Response{
		Status: fiber.StatusOK,
		Data:   "SignInResendCode",
	}

	// Se almacena el objeto para que el middleware lo procese
	ctx.JSON(fiber.StatusOK, customResponse)
}
