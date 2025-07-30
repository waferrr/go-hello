package main

import "github.com/gin-gonic/gin"

func main() {

	//创建服务
	ginServer := gin.Default()

	//接口
	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	ginServer.Run(":8080")

}
