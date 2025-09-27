package controllers

import (
	"accounts/internal/api/v1/emails/domain/entities"
	"accounts/internal/api/v1/emails/interface/dtos"
	"accounts/internal/common/requests"
	"accounts/internal/common/responses"
	"foundation/domain/customctx"
	"foundation/domain/logger"
	"foundation/interface/cdtos"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func (c *EmailsController) Activate(ctx *gin.Context) {

	cc := customctx.NewCustomContext(ctx.Request.Context())

	entry := logger.FromContext(ctx.Request.Context())

	entry.Info("Activate Controller")

	dto := cdtos.GetDTOWithResponse[dtos.ActivateDTO](ctx, cc)

	if dto.Error != nil {
		ctx.JSON(dto.StatusCode, dto.ToMapWithCustomContext(cc))
		return
	}

	organizationID, ok := ctx.Get("organization_id")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Organization ID not found",
		})
		return
	}

	command := dto.Data.ToCommand()
	command.OrganizationID = organizationID.(string)

	response := c.userService.Activate(cc, command)
	// Se almacena el objeto para que el middleware lo procese
	ctx.JSON(response.StatusCode, response.ToMapWithCustomContext(cc))
}

func (c *EmailsController) ActivateV2(ctx *gin.Context) {
	dto := requests.GetDTO[dtos.ActivateDTO](ctx)

	entity, err := entities.NewActivateFromJSON(dto.ToJson())

	if err != nil {
		customResponse := responses.Response{
			Status: fiber.StatusBadRequest,
			Data:   "Error al parsear el JSON",
		}

		// Se almacena el objeto para que el middleware lo procese
		ctx.JSON(fiber.StatusOK, customResponse)
		return
	}

	response := c.userService.ActivateV2(ctx.Request.Context(), entity)
	// Se almacena el objeto para que el middleware lo procese
	ctx.JSON(response.StatusCode, response.ToMap())
}
