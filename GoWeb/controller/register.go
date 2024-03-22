package controller

import (
	"web/middleware"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	router.GET("login", Login)

	auth := router.Group("/user", middleware.Auth(1))
	{
		auth.GET("", GetUser)
		auth.GET("/create", CreateUser)
	}
}
