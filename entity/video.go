package entity

import "time"

type (
	Author struct {
		ID        uint64 `gorm:"primary_key;auto_increment" json:"id,omitempty"`
		FirstName string `json:"firstname" binding:"required" gorm:"type:varchar(32)"`
		LastName  string `json:"lastname" binding:"required" gorm:"type:varchar(32)"`
		Age       int    `json:"age,omitempty" binding:"gte=1,lte=130"`
		Email     string `json:"email" binding:"required,email" gorm:"type:varchar(256)"`
	}
	Video struct {
		ID          uint64    `gorm:"primary_key;auto_increment" json:"id,omitempty"`
		Title       string    `json:"title" binding:"min=2,max=100" gorm:"type:varchar(100)" validate:"is-cool"`
		Description string    `json:"description" binding:"max=200" gorm:"type:varchar(200)"`
		URL         string    `json:"url" binding:"required,url" gorm:"type:varchar(256);UNIQUE"`
		Author      Author    `json:"author,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		AuthorID    uint64    `json:"-"`
		CreatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
		UpdatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	}
)
