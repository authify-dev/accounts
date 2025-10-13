package controllers

import (
	"accounts/internal/api/v1/pending_registrations/interface/dtos"
	"accounts/internal/common/logger"
	"accounts/internal/common/requests"

	"github.com/gin-gonic/gin"
)

func (c *PendingRegistrationsController) Create(ctx *gin.Context) {

	entry := logger.FromContext(ctx.Request.Context())

	entry.Info("Creating pending registration")

	dto := requests.GetDTO[dtos.CreatePendingRegistrationDTO](ctx)

	if dto == nil {
		entry.Error("Invalid request body")
		return
	}

	command := dto.ToCommand()

	response := c.service.Create(ctx.Request.Context(), command)

	ctx.JSON(response.StatusCode, response.ToMap())

}
