package model

type Draft struct {
	ID int `gorm:"primary_key"`
	Markdown_text string
	UserId int
}

type PostJsonForm struct {
	MarkdownText  string `json:"markdown_text"`
}