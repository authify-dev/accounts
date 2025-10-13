package controllers

import (
	"accounts/internal/common/logger"
	"accounts/internal/common/requests"

	"github.com/gin-gonic/gin"
)

func (c *OAuthController) UserInfoGoogle(ctx *gin.Context) {

	entry := logger.FromContext(ctx.Request.Context())
	entry.Info("UserInfoGoogle")

	token := requests.GetToken(ctx)
	if token == nil {
		entry.Error("Error obtaining token")
		return
	}

	response := c.service.GetUserInfoGoogle(ctx.Request.Context(), token.Token)
	if response.Err != nil {
		entry.Error("UserInfoGoogle", response.Err)
	}

	ctx.JSON(response.StatusCode, response.ToMap())
}
