package controllers

import (
	"accounts/internal/api/v1/emails/interface/dtos"
	"accounts/internal/common/logger"
	"foundation/domain/customctx"
	"foundation/interface/cdtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *EmailsController) SetPassword(ctx *gin.Context) {

	cc := customctx.NewCustomContext(ctx.Request.Context())

	entry := logger.FromContext(cc.Context())

	entry.Info("SetPassword Controller")

	dto := cdtos.GetDTOWithResponse[dtos.SetPasswordDTO](ctx, cc)

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

	token := cdtos.GetAuthTokenWithEarlyResponse(ctx, cc)
	if token.Err != nil {
		return
	}

	response := c.userService.SetPassword(cc, command, token.Data)
	// Se almacena el objeto para que el middleware lo procese
	ctx.JSON(response.StatusCode, response.ToMapWithCustomContext(cc))
}
