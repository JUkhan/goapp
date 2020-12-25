package setup

import (
	"io"
	"os"

	"github.com/JUkhan/goapp/db"

	"github.com/JUkhan/goapp/controller"
	"github.com/JUkhan/goapp/middleware"
	"github.com/JUkhan/goapp/repository"
	"github.com/JUkhan/goapp/routes"
	"github.com/JUkhan/goapp/service"
	"github.com/gin-gonic/gin"
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
func Init(server *gin.RouterGroup) {
	setupLogOutput()
	db.InitDB()
	videoController := controller.NewVideoController(service.NewVideoService(repository.NewViderRepository()))
	authController := controller.NewAuthController(
		service.NewLoginService(),
		service.NewJWTService(),
	)

	authRoutes := server.Group("/auth")
	routes.SetAuthRoutes(authRoutes, authController)

	// api routes
	apiRoutes := server.Group("/videos", middleware.AuthorizeJWT)

	routes.SetVideoRoutes(apiRoutes, videoController)

	// The /view endpoints are public ( no authorization is required)
	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

}
