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
	var draftList []draft
	result := db.Find(&draftList)

	return result
}