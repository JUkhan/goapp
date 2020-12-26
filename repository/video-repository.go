package repository

import (
	"github.com/JUkhan/goapp/db"
	"github.com/JUkhan/goapp/entity"
	"gorm.io/gorm"
)

type (
	VideoRepository interface {
		Save(video *entity.Video)
		Update(video *entity.Video)
		Delete(video *entity.Video)
		FindAll() []entity.Video
	}
	repository struct {
		connection *gorm.DB
	}
)

func NewViderRepository() VideoRepository {
	return &repository{
		connection: db.Connection,
	}
}
func (db *repository) Save(video *entity.Video) {
	db.connection.Create(video)
}
func (db *repository) Update(video *entity.Video) {
	db.connection.Save(&video)
	//db.connection.Model(video).Updates(video)
}
func (db *repository) Delete(video *entity.Video) {
	db.connection.Delete(video)
}
func (db *repository) FindAll() []entity.Video {
	var videos []entity.Video
	db.connection.Preload("Author").Find(&videos)
	//db.connection.Find(&videos)
	return videos
}
