package controllers

import (
	"fmt"
	"net/http"
	"net/url"

	"accounts/internal/common/logger"
	"accounts/internal/core/settings"

	"github.com/gin-gonic/gin"
)

func (c *OAuthController) RedirectGoogle(ctx *gin.Context) {
	entry := logger.FromContext(ctx.Request.Context())
	entry.Info("RedirectGoogle")

	code := ctx.Query("code")
	response := c.service.SignInGoogle(ctx.Request.Context(), code, "admin")

	if response.Err != nil {
		entry.Error("SignInGoogle", response.Err)
		// aquí podrías devolver un error 500 o similar
	}

	// Si devolvió un 201 y tenemos el JWT en Data.JWT:
	if response.StatusCode == http.StatusCreated {
		jwtToken := response.Body.JWT
		if jwtToken != "" {
			// construimos la URL de Next.js login
			qs := url.Values{}
			qs.Set("jwt", jwtToken)
			url := settings.Settings.OAUTH_REDIRECT_URL
			redirectURL := fmt.Sprintf("%s?%s", url, qs.Encode())

			entry.Infof("🔄 Redirecting to Next.js login with JWT: %s", redirectURL)
			ctx.Redirect(http.StatusSeeOther, redirectURL)
			return
		}
	}

	// si no se cumple, devolvemos la respuesta original
	ctx.JSON(response.StatusCode, response.ToMap())
}
