package main

import (
	"os"

	"github.com/JUkhan/goapp/middleware"
	"github.com/JUkhan/goapp/setup"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

func main() {

	//gin.SetMode(gin.ReleaseMode)

	server := gin.New()

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*html")

	if gin.Mode() == gin.DebugMode {
		server.Use(gindump.Dump())
	}
	server.Use(gin.Recovery(), middleware.Logger())

	setup.Init(server)

	port := os.Getenv("PORT")

	//[AWS] Elastic Beanstalk forwards request to port 5000
	if port == "" {
		port = "3000"
	}
	server.Run(":" + port)
}
