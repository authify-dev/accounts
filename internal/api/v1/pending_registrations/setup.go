package pending_registrations

import (
	"accounts/internal/api/v1/pending_registrations/domain/services"
	"accounts/internal/api/v1/pending_registrations/interface/controllers"
	"accounts/internal/core/domain/event"
	"accounts/internal/core/infrastructure/event_bus/local"
	"accounts/internal/core/settings"

	codes_pg "accounts/internal/db/postgres/codes"
	pending_registrations_pg "accounts/internal/db/postgres/pending_registrations"
	roles_pg "accounts/internal/db/postgres/role"
	users_pg "accounts/internal/db/postgres/users"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupPendingRegistrationsModule(r *gin.Engine, db *gorm.DB) {

	// Clients

	// Repositories
	repository := pending_registrations_pg.NewPendingRegistrationsPostgresRepository(db)
	rolesRepository := roles_pg.NewRolePostgresRepository(db)
	codeRepository := codes_pg.NewCodePostgresRepository(db)
	usersRepository := users_pg.NewUserPostgresRepository(db)

	// Services
	service := services.NewPendingRegistrationsService(
		repository,
		rolesRepository,
		usersRepository,
		codeRepository,
		LocalEventBus(),
	)

	// Controllers
	controller := controllers.NewPendingRegistrationsController(service)

	// Routes

	pendingRegistrations := r.Group(settings.Settings.ROOT_PATH + "/api/v1/pending-registrations")

	pendingRegistrations.POST("", controller.Create)

}

func LocalEventBus() event.EventBus {
	return local.MockEventBus()
}
