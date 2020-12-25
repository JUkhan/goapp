package setup

import (
	"io"
	"net/http"
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
func Init(server *gin.Engine) {
	setupLogOutput()
	db.InitDB()
	videoController := controller.NewVideoController(service.NewVideoService(repository.NewViderRepository()))
	loginController := controller.NewLoginController(
		service.NewLoginService(),
		service.NewJWTService(),
	)

	server.POST("/login", func(c *gin.Context) {
		token := loginController.Login(c)
		if token != "" {
			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			c.JSON(http.StatusUnauthorized, nil)
		}
	})

	// api routes
	apiRoutes := server.Group("/api", middleware.AuthorizeJWT)
	routes.VideoRoutes(apiRoutes, videoController)

	// The /view endpoints are public ( no authorization is required)
	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

}
