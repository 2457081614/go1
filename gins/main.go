package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/testGetQueryParams", func(context *gin.Context) {
		// 获取查询参数
		param := context.Param("name")
		fmt.Printf("name is %v", param)
	})

	/**
	  安装SSH
	*/
	r.POST("/ssh/install", func(context *gin.Context) {
		// 获取查询参数
		param := context.Param("name")
		fmt.Printf("name is %v", param)
	})

	r.Run(":8888")

}
