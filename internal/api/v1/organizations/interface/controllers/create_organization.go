package organizations_controllers

import (
	organizations_use_cases "accounts/internal/context/v1/organizations/app/use_cases"

	"github.com/gin-gonic/gin"
)

type CreateOrganizationController struct {
	service *organizations_use_cases.CreateOrganizationsUseCase
}

func NewCreateOrganizationController(service *organizations_use_cases.CreateOrganizationsUseCase) *CreateOrganizationController {
	return &CreateOrganizationController{service: service}
}

func (c *CreateOrganizationController) Handle(ctx *gin.Context) {
	res := c.service.Execute(ctx)
	ctx.JSON(res.StatusCode, res)
}
