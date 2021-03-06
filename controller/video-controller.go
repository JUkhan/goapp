package controller

import (
	"net/http"
	"strconv"

	"github.com/JUkhan/goapp/entity"
	"github.com/JUkhan/goapp/service"
	"github.com/JUkhan/goapp/validators"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type VideoController interface {
	Add(*gin.Context) (entity.Video, error)
	Update(*gin.Context) (entity.Video, error)
	Delete(*gin.Context) (entity.Video, error)
	FindAll() []entity.Video
	ShowAll(*gin.Context)
}

type videoController struct {
	service service.VideoService
}

var validate *validator.Validate

func NewVideoController(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &videoController{
		service: service,
	}
}

func (c *videoController) Add(ctx *gin.Context) (entity.Video, error) {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return video, err
	}
	//custom validation
	err = validate.Struct(video)
	if err != nil {
		return video, err
	}
	c.service.Add(&video)
	return video, nil
}
func (c *videoController) Update(ctx *gin.Context) (entity.Video, error) {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return video, err
	}
	//custom validation
	err = validate.Struct(video)
	if err != nil {
		return video, err
	}
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return video, err
	}
	video.ID = id
	c.service.Update(&video)
	return video, nil
}
func (c *videoController) Delete(ctx *gin.Context) (entity.Video, error) {
	var video entity.Video
	ctx.BindJSON(&video)
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return video, err
	}
	video.ID = id
	c.service.Delete(&video)
	return video, nil
}

func (c *videoController) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *videoController) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	model := gin.H{
		"title":  "video list",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", model)
}
