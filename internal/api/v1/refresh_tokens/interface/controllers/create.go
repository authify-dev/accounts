package controllers

import (
	"accounts/internal/common/logger"
	"accounts/internal/common/requests"

	"github.com/gin-gonic/gin"
)

func (c *RefreshTokensController) Create(ctx *gin.Context) {

	entry := logger.FromContext(ctx.Request.Context())

	entry.Info("Creating new JWT")

	token := requests.GetToken(ctx)
	if token == nil {
		entry.Error("Failed to get token from request")
		return
	}

	response := c.service.Create(ctx.Request.Context(), token.Token)

	ctx.JSON(response.StatusCode, response.ToMap())
}
