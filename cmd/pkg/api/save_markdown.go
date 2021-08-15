package api

import (
	"myapp/pkg/model"

	"github.com/jinzhu/gorm"
)

type draft struct {
	ID int `json:id`
	Title string `json:title`
	Markdown_text string `json:markdown_text`
}

func GetAll(db *gorm.DB) *gorm.DB {
	var draftList []draft
	result := db.Find(&draftList)

	return result
}

func SaveMarkdown(db *gorm.DB, json model.Draft) {
	db.Create(&json)
}

func DeleteDraft(db *gorm.DB, id string) {
	db.Delete(&draft{}, id)
}