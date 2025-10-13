package controllers

import (
	"accounts/internal/api/v1/oauth_logins/interface/dtos"
	"accounts/internal/common/logger"
	"accounts/internal/common/requests"

	"github.com/gin-gonic/gin"
)

func (c *OAuthController) TokenGoogle(ctx *gin.Context) {

	entry := logger.FromContext(ctx.Request.Context())
	entry.Info("TokenGoogle")

	token := requests.GetDTO[dtos.GetTokenGoogleDTO](ctx)
	if token == nil {
		entry.Error("Error obtaining token")
		return
	}

	response := c.service.GetTokenGoogle(ctx.Request.Context(), token.Code)
	if response.Err != nil {
		entry.Error("TokenGoogle", response.Err)
	}

	ctx.JSON(response.StatusCode, response.ToMap())
}
