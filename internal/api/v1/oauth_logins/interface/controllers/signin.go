package controllers

import (
	"accounts/internal/api/v1/oauth_logins/interface/dtos"
	"accounts/internal/common/logger"
	"accounts/internal/common/requests"

	"github.com/gin-gonic/gin"
)

func (c *OAuthController) SignInGoogle(ctx *gin.Context) {

	entry := logger.FromContext(ctx.Request.Context())
	entry.Info("SignInGoogle")

	token := requests.GetDTO[dtos.SigninGoogleDTO](ctx)
	if token == nil {
		entry.Error("Error obtaining token")
		return
	}

	response := c.service.SignInGoogle(ctx.Request.Context(), token.Code, token.Role)

	if response.Err != nil {
		entry.Error("SignInGoogle", response.Err)
	}

	ctx.JSON(response.StatusCode, response.ToMap())

}
