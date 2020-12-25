package main

import (
	"github.com/JUkhan/goapp/docs" // Swagger generated files
	"github.com/JUkhan/goapp/middleware"
	"github.com/JUkhan/goapp/setup"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"

	"os"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {

	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "gin-gonic - Video API"
	docs.SwaggerInfo.Description = "Gin-gonic web framework - Youtube Video API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:3000"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	//gin.SetMode(gin.ReleaseMode)

	server := gin.New()

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*html")

	if gin.Mode() == gin.DebugMode {
		server.Use(gindump.Dump())
	}
	server.Use(gin.Recovery(), middleware.Logger())
	apiRoutes := server.Group(docs.SwaggerInfo.BasePath)

	setup.Init(apiRoutes)

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	//[AWS] Elastic Beanstalk forwards request to port 5000
	if port == "" {
		port = "3000"
	}

	server.Run(":" + port)
}
