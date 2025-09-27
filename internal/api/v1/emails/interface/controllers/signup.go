package controllers

import (
	"accounts/internal/api/v1/emails/domain/entities"
	"accounts/internal/api/v1/emails/interface/dtos"
	"accounts/internal/common/requests"
	"accounts/internal/common/responses"
	"foundation/domain/customctx"
	"foundation/interface/cdtos"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (c *EmailsController) SignUp(ctx *gin.Context) {

	cc := customctx.NewCustomContext(ctx.Request.Context())

	dto := cdtos.GetDTOWithResponse[dtos.SignUpDTO](ctx, cc)

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

	if dto.Data.UserName == "" {
		dto.Data.UserName = "User_" + uuid.New().String()
	}

	dto.Data.OrganizationID = organizationID.(string)

	entity, err := entities.NewSingUpFromJSON(dto.Data.ToJson())

	if err != nil {
		customResponse := responses.Response{
			Status: fiber.StatusBadRequest,
			Data:   "Error al parsear el JSON",
		}

		// Se almacena el objeto para que el middleware lo procese
		ctx.JSON(fiber.StatusOK, customResponse)
		return
	}

	response := c.userService.SignUp(ctx.Request.Context(), entity)
	// Se almacena el objeto para que el middleware lo procese
	ctx.JSON(response.StatusCode, response.ToMap())
}

func (c *EmailsController) SignUpResendCode(ctx *gin.Context) {
	dto := requests.GetDTO[dtos.ResendActivationCodeDTO](ctx)

	entity, err := entities.NewResendActivationCodeFromJSON(dto.ToJson())

	if err != nil {
		customResponse := responses.Response{
			Status: fiber.StatusBadRequest,
			Data:   "Error al parsear el JSON",
		}

		// Se almacena el objeto para que el middleware lo procese
		ctx.JSON(fiber.StatusOK, customResponse)
		return
	}

	response := c.userService.ResendActivationCode(ctx.Request.Context(), entity)
	// Se almacena el objeto para que el middleware lo procese
	ctx.JSON(response.StatusCode, response.ToMap())
}
