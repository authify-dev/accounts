package users

import (
	"accounts/internal/api/v1/users/domain/services"
	"accounts/internal/api/v1/users/interface/controllers"
	"accounts/internal/core/settings"
	roles "accounts/internal/db/postgres/role"
	users "accounts/internal/db/postgres/users"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUsersModule(app *gin.Engine, db *gorm.DB) {

	service := services.NewUsersService(
		users.NewUserPostgresRepository(db),
		roles.NewRolePostgresRepository(db),
	)

	controller := controllers.NewUsersController(*service)

	// Rutas de users
	group := app.Group(settings.Settings.ROOT_PATH + "/api/v1/users")

	group.POST("", controller.Create)
	group.GET("", controller.List)

}
