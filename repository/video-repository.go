package repository

import (
	"fmt"

	"github.com/JUkhan/goapp/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type (
	VideoRepository interface {
		Save(video *entity.Video)
		Update(video *entity.Video)
		Delete(video *entity.Video)
		FindAll() []entity.Video
		CloseDB()
	}
	database struct {
		connection *gorm.DB
	}
)

func NewVideoRepository() VideoRepository {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&entity.Video{}, &entity.Authore{})
	return &database{
		connection: db,
	}
}
func (db *database) Save(video *entity.Video) {
	db.connection.Create(video)
	fmt.Println("id:: ", video.ID)
}
func (db *database) Update(video *entity.Video) {
	db.connection.Save(&video)
	//db.connection.Model(video).Updates(video)
}
func (db *database) Delete(video *entity.Video) {
	db.connection.Delete(video)
}
func (db *database) FindAll() []entity.Video {
	var videos []entity.Video
	db.connection.Preload("Authore").Find(&videos)
	//db.connection.Find(&videos)
	return videos
}
func (db *database) CloseDB() {
	//db.connection.
}
