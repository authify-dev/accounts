package policies

import (
	"accounts/internal/api/middlewares"
	"accounts/internal/api/v1/policies/domain/services"
	"accounts/internal/api/v1/policies/interface/controllers"
	verifier "accounts/internal/common/controllers"
	usecases "accounts/internal/context/v1/api_keys/app/use_cases"
	"accounts/internal/core/settings"
	api_keys_pg "accounts/internal/db/postgres/api_keys"
	policies_pg "accounts/internal/db/postgres/policies"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupPoliciesModule(r *gin.Engine, db *gorm.DB) {

	// clients
	verifier := verifier.NewArgon2idCrypto([]byte("pepper")) // TODO: get from env

	// repositories
	policies_repository := policies_pg.NewPoliciesPostgresRepository(db)
	api_keys_repository := api_keys_pg.NewAPIKeyPostgresRepository(db)

	// services
	policies_service := services.NewPoliciesService(policies_repository)

	// controllers
	policies_controller := controllers.NewPoliciesController(policies_service)

	// routes

	policies_route := r.Group(settings.Settings.ROOT_PATH + "/api/v1/policies")

	policies_route.Use(middlewares.APIKeyAuthMiddleware(*usecases.NewValidateAPIKeyUseCase(api_keys_repository, verifier)))

	policies_route.POST("", policies_controller.Create)
	policies_route.GET("", policies_controller.List)
}
