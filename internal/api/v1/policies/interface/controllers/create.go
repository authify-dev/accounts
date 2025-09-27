package controllers

import (
	"accounts/internal/api/v1/policies/interface/dtos"
	"accounts/internal/common/logger"
	"foundation/domain/customctx"
	"foundation/interface/cdtos"

	"github.com/gin-gonic/gin"
)

func (c *PoliciesController) Create(ctx *gin.Context) {

	entry := logger.FromContext(ctx)

	entry.Info("Creating policy")

	cc := customctx.NewCustomContext(ctx.Request.Context())

	dto := cdtos.GetDTOWithResponse[dtos.CreatePolicyDTO](ctx, cc)
	if dto.Error != nil {
		ctx.JSON(dto.StatusCode, dto.ToMapWithCustomContext(cc))
		return
	}

	policy := c.policies_service.Create(cc, dto.Data.ToCommand())

	if policy.Error != nil {
		ctx.JSON(policy.StatusCode, policy.ToMapWithCustomContext(cc))
		return
	}

	ctx.JSON(policy.StatusCode, policy.ToMapWithCustomContext(cc))
}
