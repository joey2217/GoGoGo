package main

import (
	"web/cache"
	"web/controller"
	"web/dao"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	dao.Connect()
	cache.NewClient()
	router := gin.Default()
	router.Use(dao.DaoHandler())

	controller.Register(router)

	router.Static("/static", "./static")
	router.StaticFile("/favicon.ico", "./static/favicon.ico")

	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
