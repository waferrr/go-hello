package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"laiwh/hello/model"
	"log"
	"net/http"
)

func main() {

	// 加载配置文件
	cfg, err := LoadConfig("config")
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	//创建服务
	ginServer := gin.Default()

	// 初始化数据库连接
	db, err := gorm.Open(mysql.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	// 路由：GET /users
	ginServer.GET("/users", func(c *gin.Context) {
		var users []model.User
		result := db.Find(&users)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	})
	//接口
	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	ginServer.Run(cfg.Server.Port)
}
