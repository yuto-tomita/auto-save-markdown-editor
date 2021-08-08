package model

type Draft struct {
	ID int `gorm:"primary_key"`
	Title string
	Markdown_text string
}

type PostJsonForm struct {
	Title string `json:"title"`
	MarkdownText  string `json:"markdown_text"`
}