package controllers

import (
	"accounts/internal/api/v1/api_keys/interface/dtos"
	"accounts/internal/common/logger"
	usecases "accounts/internal/context/v1/api_keys/app/use_cases"
	"foundation/domain/customctx"
	"foundation/interface/cdtos"

	"github.com/gin-gonic/gin"
)

type GeneratorAPIKeysController struct {
	service *usecases.GenerateAPIKeysUseCase
}

func NewGeneratorAPIKeysController(service *usecases.GenerateAPIKeysUseCase) *GeneratorAPIKeysController {
	return &GeneratorAPIKeysController{service: service}
}

func (c *GeneratorAPIKeysController) Handle(ctx *gin.Context) {
	cc := customctx.NewCustomContext(ctx.Request.Context())
	entry := logger.FromContext(cc.Context())

	entry.Info("Generating API key")

	dto := cdtos.GetDTOWithResponse[dtos.GenerateAPIKeysDTO](ctx, cc)

	if dto.Error != nil {
		ctx.JSON(dto.StatusCode, dto.ToMapWithCustomContext(cc))
		return
	}

	response := c.service.Generate(cc, dto.Data.ToCommand())

	ctx.JSON(response.StatusCode, response.ToMapWithCustomContext(cc))
}
