package oauth_logins

import (
	"accounts/internal/api/v1/oauth_logins/domain/services"
	"accounts/internal/api/v1/oauth_logins/interface/controllers"
	"accounts/internal/core/infrastructure/event_bus/local"
	"accounts/internal/core/settings"
	"accounts/internal/infrastucture/oauth/google/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	login_methods "accounts/internal/db/postgres/login_methods"
	oauths "accounts/internal/db/postgres/oauth_logins"
	refresh "accounts/internal/db/postgres/refresh_tokens"
	roles "accounts/internal/db/postgres/role"
	users "accounts/internal/db/postgres/users"

	utils_controller "accounts/internal/common/controllers"
)

func SetupOAuthModule(r *gin.Engine, db *gorm.DB) {

	// repositories
	google_repository := repositories.NewOAuthGoogleRepository(
		settings.Settings.GOOGLE_OAUTH_CLIENT_ID,
		settings.Settings.GOOGLE_OAUTH_CLIENT_SECRET,
		settings.Settings.GOOGLE_OAUTH_REDIRECT_URI,
	)

	oauth_repository := oauths.NewOAuthLoginPostgresRepository(db)
	user_repository := users.NewUserPostgresRepository(db)
	role_repository := roles.NewRolePostgresRepository(db)
	login_method_repository := login_methods.NewLoginMethodPostgresRepository(db)
	refresh_repository := refresh.NewRefreshTokenPostgresRepository(db)

	// services
	serv := services.NewOAuthService(
		google_repository,
		oauth_repository,
		user_repository,
		role_repository,
		login_method_repository,
		refresh_repository,
		utils_controller.JWTController{
			PublicKey:  settings.Settings.PUBLIC_KEY_JWT,
			PrivateKey: settings.Settings.PRIVATE_KEY_JWT,
		},
		local.MockEventBus(),
	)

	// controllers
	cntr := controllers.NewOAuthController(serv)

	// routes
	oauths := r.Group(settings.Settings.ROOT_PATH + "/api/v1/platforms")

	oauths.GET("/link/google", cntr.LinkGoogle)
	oauths.POST("/token/google", cntr.TokenGoogle)
	oauths.GET("/user-info/google", cntr.UserInfoGoogle)
	oauths.POST("/signin/google", cntr.SignInGoogle)
	oauths.GET("/redirect/google", cntr.RedirectGoogle)

}
