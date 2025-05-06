package oauth_logins

import (
	"accounts/internal/api/v1/oauth_logins/domain/services"
	"accounts/internal/api/v1/oauth_logins/interface/controllers"

	"github.com/gin-gonic/gin"
)

func SetupOAuthModule(r *gin.Engine) {
	// infrastructure

	// repositories

	// services
	serv := services.NewOAuthService()

	// controllers
	cntr := controllers.NewOAuthController(serv)

	// routes
	oauths := r.Group("/api/v1/platforms")

	oauths.GET("/link/google", cntr.LinkGoogle)
}
