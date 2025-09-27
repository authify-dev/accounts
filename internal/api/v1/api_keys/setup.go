package apikeys

import (
	"accounts/internal/api/middlewares"
	"accounts/internal/api/v1/api_keys/interface/controllers"
	common_controllers "accounts/internal/common/controllers"
	usecases "accounts/internal/context/v1/api_keys/app/use_cases"
	"accounts/internal/core/settings"
	apikeys_gorm "accounts/internal/db/postgres/api_keys"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAPIKeysModule(router *gin.Engine, db *gorm.DB) {

	// Clients
	generatorAPIKey := common_controllers.NewGeneratorAPIKey(48, "live")
	verifier := common_controllers.NewArgon2idCrypto([]byte("pepper")) // TODO: get from env

	// Repositories
	apiKeyRepository := apikeys_gorm.NewAPIKeyPostgresRepository(db)

	// Services
	generateAPIKeysUseCase := usecases.NewGenerateAPIKeysUseCase(generatorAPIKey, apiKeyRepository, verifier)

	// Controllers

	generator_api_keys_controller := controllers.NewGeneratorAPIKeysController(generateAPIKeysUseCase)

	api_keys_route := router.Group(settings.Settings.ROOT_PATH + "/api/v1/api-keys")
	api_keys_route.POST("/generate", generator_api_keys_controller.Handle)
	api_keys_route.GET("/validate-secret", middlewares.APIKeyAuthMiddleware(*usecases.NewValidateAPIKeyUseCase(apiKeyRepository, verifier)), simpleTest)
}

func simpleTest(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"message": "The secret key is valid",
		},
		"success": true,
		"status":  http.StatusOK,
	})
}
