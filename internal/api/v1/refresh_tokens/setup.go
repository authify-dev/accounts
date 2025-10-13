package refreshtokens

import (
	"accounts/internal/api/v1/refresh_tokens/domain/services"
	"accounts/internal/api/v1/refresh_tokens/interface/controllers"
	jwt_controller "accounts/internal/common/controllers"
	"accounts/internal/core/settings"
	postgres_login_methods "accounts/internal/db/postgres/login_methods"
	postgres_refresh_tokens "accounts/internal/db/postgres/refresh_tokens"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRefreshTokensModule(router *gin.Engine, db *gorm.DB) {

	jwt_controller := jwt_controller.JWTController{
		PublicKey:  settings.Settings.PUBLIC_KEY_JWT,
		PrivateKey: settings.Settings.PRIVATE_KEY_JWT,
	}

	// repositories
	refreshTokensRepository := postgres_refresh_tokens.NewRefreshTokenPostgresRepository(db)

	login_methods_repository := postgres_login_methods.NewLoginMethodPostgresRepository(db)

	// services
	service := services.NewRefreshTokensService(
		refreshTokensRepository,
		login_methods_repository,
		jwt_controller,
	)

	// controller
	controller := controllers.NewRefreshTokensController(service)

	refresh := router.Group(settings.Settings.ROOT_PATH + "/api/v1/refresh-jwt")

	refresh.GET("", controller.Create)

	jwt := router.Group("/api/v1/validate-jwt")
	jwt.GET("", controller.Validate)

}
