package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {

	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.POST("/ssh/install", func(c *gin.Context) {
		params := NoPasswordSsh{}
		err := c.BindJSON(&params)
		if err != nil {
			fmt.Println("convert params error", err)
		}
		fmt.Printf("params is %v \n", params)
	})

	router.Run(":9999")
}

type NoPasswordSsh struct {
	PublicKeyContent string `json:"publicKeyContent"`
}

type User struct {
	Name     string `json:"name"`
	Password int64  `json:"password"`
}
