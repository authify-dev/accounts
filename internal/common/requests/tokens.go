package requests

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthToken struct {
	Token string `json:"token" binding:"required"`
}

func (t *AuthToken) Validate() error {
	if t.Token == "" {
		return fmt.Errorf("token is required")
	}
	return nil
}

func GetToken(ctx *gin.Context) *AuthToken {
	var token AuthToken

	token_str := ctx.GetHeader("Authorization")

	if token_str == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header is required"})
		return nil
	}

	if len(token_str) < 7 || token_str[:7] != "Bearer " {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header must start with 'Bearer '"})
		return nil
	}

	token_str = token_str[len("Bearer "):]
	token.Token = token_str
	return &token
}
