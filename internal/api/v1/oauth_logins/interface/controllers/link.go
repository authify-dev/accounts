package controllers

import (
	"accounts/internal/common/logger"

	"github.com/gin-gonic/gin"
)

func (c *OAuthController) LinkGoogle(ctx *gin.Context) {

	entry := logger.FromContext(ctx.Request.Context())
	entry.Info("LinkGoogle")

	response := c.service.GetLinkGoogle(ctx.Request.Context())
	if response.Err != nil {
		entry.Error("LinkGoogle", response.Err)
	}

	ctx.JSON(response.StatusCode, response.ToMap())
}
