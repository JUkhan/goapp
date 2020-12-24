package entity

import "gorm.io/gorm"

type (
	Authore struct {
		gorm.Model
		Name string `json:"name" binding:"required,min=5"`
		Age  int8   `json:"age" binding:"required,gte=20,lte=130"`
	}
	Video struct {
		gorm.Model
		Title       string `json:"title" binding:"required,min=10,max=120" validate:"is-cool"`
		Description string `json:"description" binding:"required,min=10,max=120"`
		URL         string `json:"url" binding:"required,url`
		AuthoreID   uint64
		Authore     Authore `json:"authore"`
	}
)
