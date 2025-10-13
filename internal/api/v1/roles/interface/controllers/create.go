package controllers

import (
	"accounts/internal/api/v1/roles/interface/dtos"
	"foundation/domain/customctx"
	"foundation/interface/cdtos"
	"foundation/utils"
	"foundation/utils/cerrs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *RolesController) Create(ctx *gin.Context) {

	cc := customctx.NewCustomContext(ctx.Request.Context())

	organization_id, ok := ctx.Get("organization_id")
	if !ok {
		res := utils.Response[string]{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Error:      cc.NewError(cerrs.NewCustomError(http.StatusBadRequest, "organization_id is required", "organization_id_is_required")),
		}
		ctx.JSON(res.StatusCode, res.ToMapWithCustomContext(cc))
		return
	}

	dto := cdtos.GetDTOWithResponse[dtos.CreateRoleDTO](ctx, cc)
	if dto.Error != nil {
		ctx.JSON(dto.StatusCode, dto.ToMapWithCustomContext(cc))
		return
	}

	command := dto.Data.ToCommand()
	command.OrganizationID = organization_id.(string)

	res := c.roles_service.Create(cc, command)

	if res.Error != nil {
		ctx.JSON(res.StatusCode, res.ToMapWithCustomContext(cc))
		return
	}

	ctx.JSON(res.StatusCode, res.ToMapWithCustomContext(cc))

}
