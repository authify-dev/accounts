package oauth_logins

import (
	"accounts/internal/api/v1/oauth_logins/domain/services"
	"accounts/internal/api/v1/oauth_logins/interface/controllers"
	"accounts/internal/core/settings"
	"accounts/internal/infrastucture/oauth/google/repositories"

	"github.com/gin-gonic/gin"
)

func SetupOAuthModule(r *gin.Engine) {
	// infrastructure

	// repositories
	google_repository := repositories.NewOAuthGoogleRepository(
		settings.Settings.GOOGLE_OAUTH_CLIENT_ID,
		settings.Settings.GOOGLE_OAUTH_CLIENT_SECRET,
		settings.Settings.GOOGLE_OAUTH_REDIRECT_URI,
	)

	// services
	serv := services.NewOAuthService(
		google_repository,
	)

	// controllers
	cntr := controllers.NewOAuthController(serv)

	// routes
	oauths := r.Group("/api/v1/platforms")

	oauths.GET("/link/google", cntr.LinkGoogle)
	oauths.POST("/token/google", cntr.TokenGoogle)
	oauths.GET("/user-info/google", cntr.UserInfoGoogle)
	oauths.GET("/signin/google", cntr.UserInfoGoogle)

}
