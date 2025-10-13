package controllers

import (
	"accounts/internal/api/v1/emails/interface/dtos"
	"accounts/internal/common/logger"
	"foundation/domain/customctx"
	"foundation/interface/cdtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *EmailsController) ResetPassword(ctx *gin.Context) {

	cc := customctx.NewCustomContext(ctx.Request.Context())

	entry := logger.FromContext(cc.Context())

	entry.Info("ResetPassword")

	dto := cdtos.GetDTOWithResponse[dtos.ResetPasswordDTO](ctx, cc)

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

	response := c.userService.ResetPassword(cc, command)

	ctx.JSON(response.StatusCode, response.ToMapWithCustomContext(cc))
}

func (c *EmailsController) ResetPasswordConfirm(ctx *gin.Context) {

	cc := customctx.NewCustomContext(ctx.Request.Context())
	entry := logger.FromContext(cc.Context())

	entry.Info("ResetPassword Confirm")

	dto := cdtos.GetDTOWithResponse[dtos.ConfirmPasswordDTO](ctx, cc)

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

	response := c.userService.ConfirmPassword(cc, command)

	ctx.JSON(response.StatusCode, response.ToMapWithCustomContext(cc))
}
