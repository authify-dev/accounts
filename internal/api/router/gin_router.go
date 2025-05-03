package router

import (
	"accounts/internal/common/middlewares"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.RequestLogMiddleware())

	return r
}
