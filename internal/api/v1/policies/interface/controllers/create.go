package controllers

import (
	"accounts/internal/api/v1/policies/interface/dtos"
	"accounts/internal/common/logger"
	"foundation/domain/customctx"
	"foundation/interface/cdtos"
	"foundation/utils"
	"foundation/utils/cerrs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *PoliciesController) Create(ctx *gin.Context) {

	entry := logger.FromContext(ctx)

	entry.Info("Creating policy")

	cc := customctx.NewCustomContext(ctx.Request.Context())

	organization_id, ok := ctx.Get("organization_id")
	if !ok {
		res := utils.Response[string]{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Error:      cc.NewError(cerrs.NewCustomError(http.StatusBadRequest, "organization_id is required", "organization_id_is_required")),
		}
		ctx.JSON(
			res.StatusCode,
			res.ToMapWithCustomContext(cc),
		)
		return
	}

	dto := cdtos.GetDTOWithResponse[dtos.CreatePolicyDTO](ctx, cc)
	if dto.Error != nil {
		ctx.JSON(dto.StatusCode, dto.ToMapWithCustomContext(cc))
		return
	}

	command := dto.Data.ToCommand()
	command.OrganizationID = organization_id.(string)

	policy := c.policies_service.Create(cc, command)

	if policy.Error != nil {
		ctx.JSON(policy.StatusCode, policy.ToMapWithCustomContext(cc))
		return
	}

	ctx.JSON(policy.StatusCode, policy.ToMapWithCustomContext(cc))
}
