package controllers

import (
	"accounts/internal/common/logger"
	"foundation/domain/customctx"

	"github.com/gin-gonic/gin"
)

func (c *RolePoliciesController) Info(ctx *gin.Context) {

	cc := customctx.NewCustomContext(ctx.Request.Context())

	entry := logger.FromContext(cc.Context())

	entry.Info("Getting role policies info")

	role_id := ctx.Param("role_id")

	role_policies := c.service.Info(cc, role_id)

	ctx.JSON(role_policies.StatusCode, role_policies.ToMapWithCustomContext(cc))
}
