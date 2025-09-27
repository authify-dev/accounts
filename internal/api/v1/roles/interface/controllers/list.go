package controllers

import (
	"accounts/internal/api/v1/roles/domain/entities"
	"accounts/internal/common/logger"
	"foundation/domain/customctx"
	"foundation/utils"
	"foundation/utils/cerrs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *RolesController) List(ctx *gin.Context) {

	cc := customctx.NewCustomContext(ctx.Request.Context())

	entry := logger.FromContext(cc.Context())

	entry.Info("Listing roles")

	organizationID, ok := ctx.Get("organization_id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, utils.Response[entities.Role]{
			StatusCode: http.StatusBadRequest,
			Error: cc.NewError(
				cerrs.NewCustomError(
					http.StatusBadRequest,
					"organization_id is required",
					"roles.list.error_listing_roles",
				),
			),
			Success: false,
		})
		return
	}

	res := c.roles_service.List(cc, organizationID.(string))
	if res.Error != nil {
		ctx.JSON(res.StatusCode, res.ToMapWithCustomContext(cc))
		return
	}

	ctx.JSON(res.StatusCode, res.ToMapWithCustomContext(cc))
}
