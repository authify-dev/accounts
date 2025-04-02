package requests

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DTO interface {
	Validate() error
}

func GetDTO[K DTO](ctx *gin.Context) *K {
	var dto K
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}
	if err := dto.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}
	return &dto
}
