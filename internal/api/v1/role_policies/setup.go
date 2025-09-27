package role_policies

import (
	"accounts/internal/api/middlewares"
	"accounts/internal/api/v1/role_policies/domain/services"
	"accounts/internal/api/v1/role_policies/interface/controllers"
	verifier "accounts/internal/common/controllers"
	usecases "accounts/internal/context/v1/api_keys/app/use_cases"
	"accounts/internal/core/settings"
	api_keys_pg "accounts/internal/db/postgres/api_keys"
	policies_pg "accounts/internal/db/postgres/policies"
	roles_pg "accounts/internal/db/postgres/role"
	role_policies_pg "accounts/internal/db/postgres/role_policies"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRolePoliciesModule(router *gin.Engine, db *gorm.DB) {
	// clients
	verifier := verifier.NewArgon2idCrypto([]byte("pepper")) // TODO: get from env

	// repositories
	role_policies_repository := role_policies_pg.NewRolePoliciesPostgresRepository(db)
	role_repository := roles_pg.NewRolePostgresRepository(db)
	policies_repository := policies_pg.NewPoliciesPostgresRepository(db)
	api_keys_repository := api_keys_pg.NewAPIKeyPostgresRepository(db)
	// services
	role_policies_service := services.NewRolePoliciesService(role_policies_repository, role_repository, policies_repository)

	// controllers
	role_policies_controller := controllers.NewRolePoliciesController(role_policies_service)

	// routes

	role_policies_route := router.Group(settings.Settings.ROOT_PATH + "/api/v1/role_policies")

	role_policies_route.Use(middlewares.APIKeySecretAuthMiddleware(*usecases.NewValidateSecretAPIKeyUseCase(api_keys_repository, verifier)))

	role_policies_route.POST("", role_policies_controller.Create)
	role_policies_route.GET("/:role_id", role_policies_controller.Info)
}
