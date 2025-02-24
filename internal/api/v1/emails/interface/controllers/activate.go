package controllers

import (
	"accounts/internal/common/responses"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func (c *EmailsController) Activate(ctx *gin.Context) {
	customResponse := responses.Response{
		Status: fiber.StatusOK,
		Data:   "Activate",
	}

	// Se almacena el objeto para que el middleware lo procese
	ctx.JSON(fiber.StatusOK, customResponse)
}
