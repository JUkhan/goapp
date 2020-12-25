package routes

import (
	"github.com/JUkhan/goapp/api"
	"github.com/JUkhan/goapp/controller"
	"github.com/gin-gonic/gin"
)

func SetAuthRoutes(router *gin.RouterGroup, authController controller.AuthController) {
	handler := api.NewAuthAPI(authController)
	router.POST("/token", handler.Authenticate)
}
