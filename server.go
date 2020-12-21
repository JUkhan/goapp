package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()
	aVideo := entity.Video{}
	server.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	server.Run(":3000")
}
