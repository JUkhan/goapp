package routes

import (
	"net/http"

	"github.com/JUkhan/goapp/controller"
	"github.com/gin-gonic/gin"
)

func VideoRoutes(router *gin.RouterGroup, videoController controller.VideoController) {

	router.GET("/videos", func(c *gin.Context) {
		c.JSON(200, videoController.FindAll())
	})
	router.POST("/videos", func(c *gin.Context) {
		v, err := videoController.Add(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(200, v)
		}
	})
	router.PUT("/videos/:id", func(c *gin.Context) {
		v, err := videoController.Update(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(200, v)
		}
	})
	router.DELETE("/videos/:id", func(c *gin.Context) {
		v, err := videoController.Delete(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(200, v)
		}
	})
}
