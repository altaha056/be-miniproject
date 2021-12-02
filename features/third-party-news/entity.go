package news

import "time"

type NewsCore struct {
	Author      string    `json: "original_title"`
	Category    string    `json: "overview"`
	Country     string    `json: "country"`
	Description string    `json: "description"`
	Image       string    `json: "image"`
	Language    string    `json: "language"`
	PublishedAt time.Time `json: "published_at"`
	Source      string    `json: "source"`
	Title       string    `json: "title"`
	Url         string    `json: "url"`
}

type Business interface {
	GetNews(keyword string) ([]NewsCore, error)
}

type Data interface {
	GetData(keyword string) ([]NewsCore, error)
}
