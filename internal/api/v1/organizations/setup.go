package organizations

import (
	organizations_controllers "accounts/internal/api/v1/organizations/interface/controllers"
	organizations_use_cases "accounts/internal/context/v1/organizations/app/use_cases"
	"accounts/internal/core/settings"
	organizations_gorm "accounts/internal/db/postgres/organinizations"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupOrganizationsModule(r *gin.Engine, db *gorm.DB) {

	// Clients

	// Repositories
	organizationsRepository := organizations_gorm.NewOrganizationsPostgresRepository(db)

	// Use Cases
	createOrganizationsUseCase := organizations_use_cases.NewCreateOrganizationsUseCase(organizationsRepository)

	// Controllers
	createOrganizationController := organizations_controllers.NewCreateOrganizationController(createOrganizationsUseCase)

	organizationsRouter := r.Group(settings.Settings.ROOT_PATH + "/v1/organizations")

	organizationsRouter.POST("", createOrganizationController.Handle)
}
