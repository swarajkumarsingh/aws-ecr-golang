package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "hello world")
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, "health ok")
	})

	r.Run(":8080")
}
