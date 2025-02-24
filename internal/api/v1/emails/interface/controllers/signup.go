package controllers

import (
	"accounts/internal/common/responses"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func (c *EmailsController) SignUp(ctx *gin.Context) {
	customResponse := responses.Response{
		Status: fiber.StatusOK,
		Data:   "SignUp",
	}

	// Se almacena el objeto para que el middleware lo procese
	ctx.JSON(fiber.StatusOK, customResponse)
}

func (c *EmailsController) SignUpResendCode(ctx *gin.Context) {
	customResponse := responses.Response{
		Status: fiber.StatusOK,
		Data:   "SignUpResendCode",
	}

	// Se almacena el objeto para que el middleware lo procese
	ctx.JSON(fiber.StatusOK, customResponse)
}
