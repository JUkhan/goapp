package routes

import (
	"github.com/JUkhan/goapp/api"
	"github.com/JUkhan/goapp/controller"
	"github.com/gin-gonic/gin"
)

func SetVideoRoutes(router *gin.RouterGroup, videoController controller.VideoController) {
	handler := api.NewVideoAPI(videoController)
	router.GET("", handler.GetVideos)
	router.POST("", handler.CreateVideo)
	router.PUT(":id", handler.UpdateVideo)
	router.DELETE(":id", handler.RemoveVideo)
}
