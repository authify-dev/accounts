package middlewares

import (
	usecases "accounts/internal/context/v1/api_keys/app/use_cases"
	"foundation/domain/customctx"
	"net/http"

	"github.com/gin-gonic/gin"
)

func APIKeyAuthMiddleware(usecases usecases.ValidateAPIKeyUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-KEY")
		if apiKey == "" {
			c.JSON(
				http.StatusUnauthorized, gin.H{
					"error":   "API key is required",
					"success": false,
					"status":  http.StatusUnauthorized,
				},
			)
			c.Abort()
			return
		}

		cc := customctx.NewCustomContext(c.Request.Context())

		res := usecases.Validate(cc, apiKey)

		if res.Error != nil {
			c.JSON(res.StatusCode, res.ToMapWithCustomContext(cc))
			c.Abort()
			return
		}

		c.Set("organization_id", res.Data.OrganizationID)

		c.Next()
	}
}
