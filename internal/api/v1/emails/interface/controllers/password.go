package controllers

import (
	"accounts/internal/api/v1/emails/interface/dtos"
	"accounts/internal/common/logger"
	"accounts/internal/common/requests"

	"github.com/gin-gonic/gin"
)

func (c *EmailsController) ResetPassword(ctx *gin.Context) {

	entry := logger.FromContext(ctx.Request.Context())

	entry.Info("ResetPassword")

	dto := requests.GetDTO[dtos.ResetPasswordDTO](ctx)

	entity := dto.ToEntity()
	entry.Infof("DTO: %v", dto)
	entry.Infof("Entity: %v", entity)

	response := c.userService.ResetPassword(ctx.Request.Context(), entity)

	ctx.JSON(response.StatusCode, response.ToMap())
}

func (c *EmailsController) ResetPasswordConfirm(ctx *gin.Context) {
	entry := logger.FromContext(ctx.Request.Context())

	entry.Info("ResetPassword Confirm")

	dto := requests.GetDTO[dtos.ConfirmPasswordDTO](ctx)

	entity := dto.ToEntity()
	entry.Infof("DTO: %v", dto)
	entry.Infof("Entity: %v", entity)

	response := c.userService.ConfirmPassword(ctx.Request.Context(), entity)

	ctx.JSON(response.StatusCode, response.ToMap())
}
