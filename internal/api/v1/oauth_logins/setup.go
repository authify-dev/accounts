package oauth_logins

import (
	"accounts/internal/api/v1/oauth_logins/domain/services"
	"accounts/internal/api/v1/oauth_logins/interface/controllers"
	"accounts/internal/infrastucture/oauth/google/repositories"

	"github.com/gin-gonic/gin"
)

func SetupOAuthModule(r *gin.Engine) {
	// infrastructure

	// repositories
	google_repository := repositories.NewOAuthGoogleRepository()

	// services
	serv := services.NewOAuthService(
		google_repository,
	)

	// controllers
	cntr := controllers.NewOAuthController(serv)

	// routes
	oauths := r.Group("/api/v1/platforms")

	oauths.GET("/link/google", cntr.LinkGoogle)
}
