package controllers

import (
	"accounts/internal/api/v1/roles/interface/dtos"
	"foundation/domain/customctx"
	"foundation/interface/cdtos"

	"github.com/gin-gonic/gin"
)

func (c *RolesController) Create(ctx *gin.Context) {

	cc := customctx.NewCustomContext(ctx.Request.Context())

	dto := cdtos.GetDTOWithResponse[dtos.CreateRoleDTO](ctx, cc)
	if dto.Error != nil {
		ctx.JSON(dto.StatusCode, dto.ToMapWithCustomContext(cc))
		return
	}

	res := c.roles_service.Create(cc, dto.Data.ToCommand())

	if res.Error != nil {
		ctx.JSON(res.StatusCode, res.ToMapWithCustomContext(cc))
		return
	}

	ctx.JSON(res.StatusCode, res.ToMapWithCustomContext(cc))

}
