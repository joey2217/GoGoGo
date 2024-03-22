package dao

import (
	"web/dao/model"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	_db *gorm.DB
)

func Connect() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:M68:jpYgA2btiPd@tcp(127.0.0.1:9999)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}
	_db = db
	migrate()
}

func migrate() {
	_db.AutoMigrate(&model.UserDO{})
}

func DaoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("DB", _db.WithContext(c))
		c.Next()
	}
}

func Repo(ctx *gin.Context) *gorm.DB {
	return ctx.MustGet("DB").(*gorm.DB)
}
