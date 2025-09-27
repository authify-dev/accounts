package apikeys

import (
	"accounts/internal/api/v1/api_keys/interface/controllers"
	api_keys_controllers "accounts/internal/common/controllers"
	usecases "accounts/internal/context/v1/api_keys/app/use_cases"
	"accounts/internal/core/settings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAPIKeysModule(router *gin.Engine, db *gorm.DB) {

	// Clients
	generatorAPIKey := api_keys_controllers.NewGeneratorAPIKey(48)

	// Services
	generateAPIKeysUseCase := usecases.NewGenerateAPIKeysUseCase(generatorAPIKey)

	// Controllers

	generator_api_keys_controller := controllers.NewGeneratorAPIKeysController(generateAPIKeysUseCase)

	api_keys_route := router.Group(settings.Settings.ROOT_PATH + "/api/v1/api-keys")
	api_keys_route.POST("/generate", generator_api_keys_controller.Handle)
}
