package model

type Book struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishedAt string `json:"published_at"`
	Content     string `json:"content"`
}
