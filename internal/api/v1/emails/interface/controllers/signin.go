package controllers

import (
	"accounts/internal/api/v1/emails/interface/dtos"
	"accounts/internal/common/responses"
	"foundation/domain/customctx"
	"foundation/domain/logger"
	"foundation/interface/cdtos"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func (c *EmailsController) SignIn(ctx *gin.Context) {

	cc := customctx.NewCustomContext(ctx.Request.Context())

	entry := logger.FromContext(cc.Context())

	entry.Info("SignIn")

	dto := cdtos.GetDTOWithResponse[dtos.SignInDTO](ctx, cc)

	if dto.Error != nil {
		ctx.JSON(dto.StatusCode, dto.ToMapWithCustomContext(cc))
		return
	}

	command := dto.Data.ToCommand()

	organizationID, ok := ctx.Get("organization_id")

	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Organization ID not found",
		})
		return
	}

	command.OrganizationID = organizationID.(string)

	response := c.userService.SignIn(cc, command)

	ctx.JSON(response.StatusCode, response.ToMapWithCustomContext(cc))
}

func (c *EmailsController) SignInResendCode(ctx *gin.Context) {
	customResponse := responses.Response{
		Status: fiber.StatusOK,
		Data:   "SignInResendCode",
	}

	// Se almacena el objeto para que el middleware lo procese
	ctx.JSON(fiber.StatusOK, customResponse)
}
