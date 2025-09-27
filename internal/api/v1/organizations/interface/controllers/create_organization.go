package organizations_controllers

import (
	organizations_use_cases "accounts/internal/context/v1/organizations/app/use_cases"
	"foundation/domain/logger"

	"github.com/gin-gonic/gin"
)

type CreateOrganizationController struct {
	service *organizations_use_cases.CreateOrganizationsUseCase
}

func NewCreateOrganizationController(service *organizations_use_cases.CreateOrganizationsUseCase) *CreateOrganizationController {
	return &CreateOrganizationController{service: service}
}

func (c *CreateOrganizationController) Handle(ctx *gin.Context) {

	entry := logger.FromContext(ctx)

	entry.Info("Creating organization")

	res := c.service.Execute(ctx)
	ctx.JSON(res.StatusCode, res)
}
