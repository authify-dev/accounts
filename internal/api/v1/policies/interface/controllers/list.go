package controllers

import (
	"accounts/internal/common/logger"
	"accounts/internal/common/responses"
	"foundation/domain/customctx"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func (c *PoliciesController) List(ctx *gin.Context) {

	cc := customctx.NewCustomContext(ctx.Request.Context())

	entry := logger.FromContext(cc.Context())

	entry.Info("Listing policies")

	organizationID := ctx.Query("organization_id")
	if organizationID == "" {
		ctx.JSON(fiber.StatusBadRequest, responses.Response{
			Status: fiber.StatusBadRequest,
			Errors: []string{"organization_id is required"},
		})
		return
	}

	policies := c.policies_service.List(cc, organizationID)
	if policies.Error != nil {
		ctx.JSON(policies.StatusCode, policies.ToMapWithCustomContext(cc))
		return
	}

	ctx.JSON(policies.StatusCode, policies.ToMapWithCustomContext(cc))
}
