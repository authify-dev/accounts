package router

import (
	"accounts/internal/common/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.RequestLogMiddleware())

	r.Use(cors.Default())

	return r
}
