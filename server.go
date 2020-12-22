package main

import (
	"io"
	"net/http"
	"os"

	"github.com/JUkhan/goapp/controller"
	"github.com/JUkhan/goapp/middleware"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

func setupLogoutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
func main() {
	setupLogoutput()
	server := gin.New()
	server.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth(), gindump.Dump())
	videoController := controller.NewVideoController()

	server.GET("/videos", func(c *gin.Context) {
		c.JSON(200, videoController.FindAll())
	})
	server.POST("/videos", func(c *gin.Context) {
		v, err := videoController.Add(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(200, v)
		}
	})

	server.Run(":3000")
}
