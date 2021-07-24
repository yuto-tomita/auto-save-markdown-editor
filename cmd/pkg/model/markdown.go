package domain

type Draft struct {
	ID int `gorm:"primary_key"`
	Markdown_text string
	UserId int
}