package controllers

import (
	"accounts/internal/common/logger"
	"accounts/internal/common/requests"

	"github.com/gin-gonic/gin"
)

func (c *RefreshTokensController) Validate(ctx *gin.Context) {

	entry := logger.FromContext(ctx.Request.Context())

	entry.Info("Validate JWT")

	token := requests.GetToken(ctx)
	if token == nil {
		entry.Error("Failed to get token from request")
		return
	}

	response := c.service.Validate(ctx.Request.Context(), token.Token)

	ctx.JSON(response.StatusCode, response.ToMap())
}
