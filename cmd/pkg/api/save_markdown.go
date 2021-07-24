package api

import (
	domain "myapp/pkg/model"

	"github.com/jinzhu/gorm"
)

func GetAll(db *gorm.DB) *gorm.DB {
	draft := domain.Draft{}
	result := db.Find(draft)
	return result
}