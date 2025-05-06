package oauth_logins

import (
	"accounts/internal/api/v1/oauth_logins/domain/services"
	"accounts/internal/api/v1/oauth_logins/interface/controllers"
	"accounts/internal/core/settings"
	"accounts/internal/infrastucture/oauth/google/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	oauths "accounts/internal/db/postgres/oauth_logins"
	//login_methods "accounts/internal/db/postgres/login_methods"
	//refresh "accounts/internal/db/postgres/refresh_tokens"
	//roles "accounts/internal/db/postgres/role"
	//users "accounts/internal/db/postgres/users"
)

func SetupOAuthModule(r *gin.Engine) {
	// infrastructure
	db, err := gorm.Open(postgres.Open(settings.Settings.POSTGRES_DSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// repositories
	google_repository := repositories.NewOAuthGoogleRepository(
		settings.Settings.GOOGLE_OAUTH_CLIENT_ID,
		settings.Settings.GOOGLE_OAUTH_CLIENT_SECRET,
		settings.Settings.GOOGLE_OAUTH_REDIRECT_URI,
	)

	oauth_repository := oauths.NewOAuthLoginPostgresRepository(db)

	// services
	serv := services.NewOAuthService(
		google_repository,
		oauth_repository,
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
