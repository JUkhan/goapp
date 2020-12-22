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

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*html")

	server.Use(gin.Recovery(), middleware.Logger(), gindump.Dump())
	videoController := controller.NewVideoController()

	apiRoutes := server.Group("/api", middleware.BasicAuth())
	{

		apiRoutes.GET("/videos", func(c *gin.Context) {
			c.JSON(200, videoController.FindAll())
		})
		apiRoutes.POST("/videos", func(c *gin.Context) {
			v, err := videoController.Add(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				c.JSON(200, v)
			}
		})
	}

	// The /view endpoints are public ( no authorization is required)
	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	port := os.Getenv("PORT")

	//[AWS] Elastic Beanstalk forwards request to port 5000
	if port == "" {
		port = "3000"
	}
	server.Run(":" + port)
}
