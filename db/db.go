package db

import (
	"log"

	"github.com/JUkhan/goapp/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("database connected")
	db.AutoMigrate(&entity.Video{}, &entity.Author{})
	log.Println("database migrated")
	Connection = db

}
