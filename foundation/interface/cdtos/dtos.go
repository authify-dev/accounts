package cdtos

import (
	"foundation/domain/customctx"
	"foundation/domain/logger"
	"foundation/utils"
	"foundation/utils/cerrs"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

type DTO interface {
	Validate() error
}

type ErrorDTO struct {
	cerrs.CustomError
}

func GetDTO[K DTO](ctx *gin.Context, cc *customctx.CustomContext) *K {

	entry := logger.FromContext(ctx.Request.Context())

	var dto K
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		entry.Error(err)

		response := makeResponseError(err, dto)

		cc.NewError(response.Error)
		ctx.JSON(response.StatusCode, response.ToMap())
		return nil
	}
	if err := dto.Validate(); err != nil {
		entry.Error(err)

		response := makeResponseError(err, dto)

		cc.NewError(response.Error)
		//ctx.JSON(response.StatusCode, response.ToMapWithCustomContext(cc))
		return nil
	}
	return &dto
}

func GetDTOWithResponse[K DTO](ctx *gin.Context, cc *customctx.CustomContext) utils.Response[K] {

	entry := logger.FromContext(ctx.Request.Context())

	var dto K
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		entry.Error(err)

		response := makeResponseError(err, dto)

		cc.NewError(response.Error)
		return response
	}
	if err := dto.Validate(); err != nil {
		entry.Error(err)

		response := makeResponseError(err, dto)

		cc.NewError(response.Error)
		//ctx.JSON(response.StatusCode, response.ToMapWithCustomContext(cc))
		return response
	}
	return utils.Response[K]{
		Data:       dto,
		StatusCode: http.StatusOK,
		Success:    true,
	}
}

func makeResponseError[K DTO](err error, dto K) utils.Response[K] {

	typ := reflect.TypeOf(dto)

	// Si es un puntero, obten el elemento apuntado
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	response := utils.Response[K]{
		Error: &cerrs.CustomError{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
			Scope:   "dto.validate." + typ.Name(),
		},
		StatusCode: http.StatusUnprocessableEntity,
		Success:    false,
	}

	return response
}

func GetAuthToken(ctx *gin.Context) utils.Result[string] {

	entry := logger.FromContext(ctx.Request.Context())

	token := ctx.GetHeader("Authorization")

	if token == "" {
		entry.Error("No se encontró un token de autenticación")
		return utils.Result[string]{
			Err: &cerrs.CustomError{
				Code:    http.StatusUnauthorized,
				Message: "Not Found Authorization Header",
				Scope:   "auth.get_token.not_found_authorization_header",
			},
		}
	}

	token = strings.TrimPrefix(token, "Bearer ")

	if token == "" {
		entry.Error("No se encontró un token de autenticación")
		return utils.Result[string]{
			Err: &cerrs.CustomError{
				Code:    http.StatusUnauthorized,
				Message: "Not Found Token in Authorization Header",
				Scope:   "auth.get_token.not_found_token",
			},
		}
	}

	return utils.Result[string]{
		Data: token,
		Err:  nil,
	}
}

func GetAuthTokenWithEarlyResponse(ctx *gin.Context, cc *customctx.CustomContext) utils.Result[string] {

	entry := logger.FromContext(ctx.Request.Context())

	token := ctx.GetHeader("Authorization")

	if token == "" {
		entry.Error("Not Found Authorization Header")

		err := cerrs.CustomError{
			Code:    http.StatusUnauthorized,
			Message: "Not Found Authorization Header",
			Scope:   "auth.get_token.not_found_authorization_header",
		}

		cc.NewError(&err)

		response := utils.Response[string]{
			StatusCode: http.StatusUnauthorized,
		}

		ctx.JSON(response.StatusCode, response.ToMapWithCustomContext(cc))
		ctx.Abort()
		return utils.Result[string]{
			Err: &err,
		}
	}

	token = strings.TrimPrefix(token, "Bearer ")

	if token == "" {
		entry.Error("Not Found Token in Authorization Header")

		err := cerrs.CustomError{
			Code:    http.StatusUnauthorized,
			Message: "Not Found Token in Authorization Header",
			Scope:   "auth.get_token.not_found_token",
		}
		cc.NewError(&err)

		response := utils.Response[string]{
			StatusCode: http.StatusUnauthorized,
		}
		ctx.JSON(response.StatusCode, response.ToMapWithCustomContext(cc))
		ctx.Abort()
		return utils.Result[string]{
			Err: &err,
		}
	}

	return utils.Result[string]{
		Data: token,
		Err:  nil,
	}
}
