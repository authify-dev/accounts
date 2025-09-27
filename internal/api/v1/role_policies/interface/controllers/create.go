package controllers

import (
	"accounts/internal/api/v1/role_policies/interface/dtos"
	"accounts/internal/common/logger"
	"accounts/internal/common/responses"
	"foundation/domain/customctx"
	"foundation/interface/cdtos"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func (c *RolePoliciesController) Create(ctx *gin.Context) {

	cc := customctx.NewCustomContext(ctx.Request.Context())
	entry := logger.FromContext(cc.Context())

	entry.Info("Creating role policies")

	organizationID := ctx.Query("organization_id")
	if organizationID == "" {
		entry.Error("organization_id is required")
		ctx.JSON(fiber.StatusBadRequest, responses.Response{
			Status: fiber.StatusBadRequest,
			Errors: []string{"organization_id is required"},
		})
		return
	}

	dto := cdtos.GetDTOWithResponse[dtos.CreateRolePoliciesDTO](ctx, cc)
	if dto.Error != nil {
		entry.Error("Invalid request")
		ctx.JSON(dto.StatusCode, dto.ToMapWithCustomContext(cc))
		return
	}

	command := dto.Data.ToCommand()

	command.OrganizationID = organizationID

	role_policies := c.service.Create(cc, command)

	ctx.JSON(role_policies.StatusCode, role_policies.ToMapWithCustomContext(cc))
}
