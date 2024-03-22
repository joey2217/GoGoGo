package controller

import (
	"web/cache"
	"web/dao"
	"web/dao/model"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	var user model.UserDO
	if err := cache.GetModel("user", &user); err == nil {
		ctx.JSON(200, user)
		return
	}
	dao.Repo(ctx).First(&user, 1)
	// 缓存
	cache.SetModel("user", user)
	ctx.JSON(200, user)
}

func CreateUser(ctx *gin.Context) {
	user := model.UserDO{
		Username: "Username",
		Password: "Username",
		Role:     0,
	}
	res := dao.Repo(ctx).Create(&user)
	if res.Error != nil {
		ctx.JSON(500, gin.H{"error": res.Error.Error()})
	} else {
		ctx.JSON(200, gin.H{"id": user.ID})
	}

}
