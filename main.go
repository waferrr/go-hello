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

	//创建服务
	ginServer := gin.Default()

	// 构造 DSN（数据源名称）
	dsn := "root:123456@tcp(47.236.236.227:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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

	ginServer.Run(":8080")

}
