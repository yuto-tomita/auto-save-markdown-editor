package model

type Draft struct {
	ID int `gorm:"primary_key"`
	Title string
	Markdown_text string
}

type PostJsonForm struct {
	MarkdownText  string `json:"markdown_text"`
}