package middlewares

import (
	usecases "accounts/internal/context/v1/api_keys/app/use_cases"
	"foundation/domain/customctx"
	"foundation/utils"
	"foundation/utils/cerrs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func APIKeyPublicAuthMiddleware(usecases usecases.ValidatePublicAPIKeyUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		cc := customctx.NewCustomContext(c.Request.Context())
		apiKey := c.GetHeader("X-API-KEY")
		if apiKey == "" {
			res := utils.Response[string]{
				Success:    false,
				StatusCode: http.StatusUnauthorized,
				Error:      cc.NewError(cerrs.NewCustomError(http.StatusUnauthorized, "API key is required", "api_key_is_required")),
			}
			c.JSON(
				res.StatusCode,
				res.ToMapWithCustomContext(cc),
			)
			c.Abort()
			return
		}

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
