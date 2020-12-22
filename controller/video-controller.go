package controller

import (
	"github.com/JUkhan/goapp/entity"
	"github.com/JUkhan/goapp/service"
	"github.com/JUkhan/goapp/validators"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type VideoController interface {
	Add(*gin.Context) (entity.Video, error)
	FindAll() []entity.Video
}

type videoController struct {
	service service.VideoService
}

var validate *validator.Validate

func NewVideoController() VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &videoController{
		service: service.NewVideoService(),
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
	c.service.Add(video)
	return video, nil
}

func (c *videoController) FindAll() []entity.Video {
	return c.service.FindAll()
}
