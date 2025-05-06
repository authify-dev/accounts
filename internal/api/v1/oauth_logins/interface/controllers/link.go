package controllers

import "github.com/gin-gonic/gin"

func (c *OAuthController) LinkGoogle(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"data":    "OAuthController",
		"success": true,
	})
}
