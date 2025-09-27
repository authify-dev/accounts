package organizations_controllers

import (
	"accounts/internal/api/v1/organizations/interface/dtos"
	organizations_use_cases "accounts/internal/context/v1/organizations/app/use_cases"
	"foundation/domain/customctx"
	"foundation/domain/logger"
	"foundation/interface/cdtos"

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

	cc := customctx.NewCustomContext(ctx.Request.Context())

	dto := cdtos.GetDTOWithResponse[dtos.CreateOrganizationDTO](ctx, cc)
	if dto.Error != nil {
		ctx.JSON(dto.StatusCode, dto.ToMapWithCustomContext(cc))
		return
	}

	res := c.service.Execute(cc, dto.Data.ToCommand())

	ctx.JSON(res.StatusCode, res.ToMapWithCustomContext(cc))
}
