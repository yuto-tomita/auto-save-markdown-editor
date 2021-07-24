package api

import (
	"github.com/jinzhu/gorm"
)

type draft struct {
	ID int `json:id`
	Markdown_text string `json:markdown_text`
	UserId int `json:user_id`
}

func GetAll(db *gorm.DB) *gorm.DB {
	result := db.Find(&draft{})
	return result
}