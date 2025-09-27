package controllers

import (
	"accounts/internal/api/v1/role_policies/interface/dtos"
	"accounts/internal/common/logger"
	"accounts/internal/common/requests"
	"accounts/internal/common/responses"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func (c *RolePoliciesController) Create(ctx *gin.Context) {
	entry := logger.FromContext(ctx)

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

	dto := requests.GetDTO[dtos.CreateRolePoliciesDTO](ctx)
	if dto == nil {
		entry.Error("Invalid request")
		return
	}

	command := dto.ToCommand()

	command.OrganizationID = organizationID

	role_policies := c.service.Create(ctx.Request.Context(), command)

	ctx.JSON(role_policies.StatusCode, role_policies.ToMap())
}
