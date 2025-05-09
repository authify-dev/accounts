package controllers

import (
	"accounts/internal/common/logger"

	"github.com/gin-gonic/gin"
)

func (c *OAuthController) RedirectGoogle(ctx *gin.Context) {

	entry := logger.FromContext(ctx.Request.Context())
	entry.Info("RedirectGoogle")

	code := ctx.Query("code")

	response := c.service.SignInGoogle(ctx.Request.Context(), code, "admin")

	if response.Err != nil {
		entry.Error("SignInGoogle", response.Err)
	}

	ctx.JSON(response.StatusCode, response.ToMap())
}
