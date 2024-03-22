package controller

import (
	"web/middleware"
	"web/utils"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	if token, err := utils.GenerateToken(1); err != nil {
		ctx.JSON(400, gin.H{
			"msg": err.Error(),
		})
	} else {
		ctx.SetCookie(middleware.AUTH_COOKIE, token, 3600, "/", "", false, true) // 设置cookie
		ctx.JSON(200, gin.H{
			"msg": "登录成功",
		})
	}
}
