package roles

import (
	"accounts/internal/api/middlewares"
	"accounts/internal/api/v1/roles/domain/services"
	"accounts/internal/api/v1/roles/interface/controllers"
	verifier "accounts/internal/common/controllers"
	usecases "accounts/internal/context/v1/api_keys/app/use_cases"
	"accounts/internal/core/settings"
	api_keys_pg "accounts/internal/db/postgres/api_keys"
	postgres_role "accounts/internal/db/postgres/role"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRolesModule(app *gin.Engine, db *gorm.DB) {

	verifier := verifier.NewArgon2idCrypto([]byte("pepper")) // TODO: get from env

	api_keys_repository := api_keys_pg.NewAPIKeyPostgresRepository(db)

	rolesService := services.NewRolesService(
		postgres_role.NewRolePostgresRepository(db),
	)

	rolesController := controllers.NewRolesController(*rolesService)

	// Rutas de users
	roles := app.Group(settings.Settings.ROOT_PATH + "/api/v1/roles")

	roles.Use(middlewares.APIKeySecretAuthMiddleware(*usecases.NewValidateSecretAPIKeyUseCase(api_keys_repository, verifier)))

	roles.POST("", rolesController.Create)
	roles.GET("", rolesController.List)

}
