package api

import (
	"net/http"

	"github.com/JUkhan/goapp/controller"
	"github.com/JUkhan/goapp/dto"
	"github.com/gin-gonic/gin"
)

type VideoApi struct {
	videoController controller.VideoController
}

func NewVideoAPI(videoController controller.VideoController) *VideoApi {
	return &VideoApi{
		videoController: videoController,
	}
}

// Paths Information

// GetVideos godoc
// @Security bearerAuth
// @Summary List existing videos
// @Description Get all the existing videos
// @Tags videos,list
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.Video
// @Failure 401 {object} dto.Response
// @Router /videos [get]
func (api *VideoApi) GetVideos(c *gin.Context) {
	c.JSON(200, api.videoController.FindAll())
}

// CreateVideo godoc
// @Security bearerAuth
// @Summary Create new videos
// @Description Create a new video
// @Tags videos,create
// @Accept  json
// @Produce  json
// @Param video body entity.Video true "Create video"
// @Success 200 {object} entity.Video
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /videos [post]
func (api *VideoApi) CreateVideo(c *gin.Context) {
	v, err := api.videoController.Add(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Response{Message: err.Error()})
	} else {
		c.JSON(200, v)
	}
}

// UpdateVideo godoc
// @Security bearerAuth
// @Summary Update videos
// @Description Update a single video
// @Security bearerAuth
// @Tags videos
// @Accept  json
// @Produce  json
// @Param  id path int true "Video ID"
// @Param video body entity.Video true "Update video"
// @Success 200 {object} entity.Video
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /videos/{id} [put]
func (api *VideoApi) UpdateVideo(c *gin.Context) {
	v, err := api.videoController.Update(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Response{Message: err.Error()})
	} else {
		c.JSON(200, v)
	}
}

// DeleteVideo godoc
// @Security bearerAuth
// @Summary Remove videos
// @Description Delete a single video
// @Security bearerAuth
// @Tags videos
// @Accept  json
// @Produce  json
// @Param  id path int true "Video ID"
// @Success 200 {object} entity.Video
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /videos/{id} [delete]
func (api *VideoApi) RemoveVideo(c *gin.Context) {
	v, err := api.videoController.Delete(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.Response{Message: err.Error()})
	} else {
		c.JSON(200, v)
	}
}
