package router

import (
	"accounts/internal/api/middlewares"
	"foundation/domain/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	r.RedirectTrailingSlash = true
	r.RedirectFixedPath = true
	r.Use(gin.Recovery())

	r.Use(middlewares.RequestLogMiddleware())
	r.Use(middlewares.TraceMiddleware())
	r.Use(middlewares.LoggerMiddleware())
	r.Use(middlewares.FlushLogsOnFinishMiddleware(logger.GetLokiHook()))

	r.Use(cors.Default())

	return r
}
