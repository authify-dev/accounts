package roles

import (
	"accounts/internal/api/v1/roles/domain/services"
	"accounts/internal/api/v1/roles/interface/controllers"
	"accounts/internal/core/settings"
	postgres_role "accounts/internal/db/postgres/role"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRolesModule(app *gin.Engine, db *gorm.DB) {

	rolesService := services.NewRolesService(
		postgres_role.NewRolePostgresRepository(db),
	)

	rolesController := controllers.NewRolesController(*rolesService)

	// Rutas de users
	roles := app.Group(settings.Settings.ROOT_PATH + "/api/v1/roles")

	roles.POST("", rolesController.Create)
	roles.GET("", rolesController.List)

}
