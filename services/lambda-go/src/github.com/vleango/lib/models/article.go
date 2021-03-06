package models

import (
	"time"
)

type Article struct {
	ID        string    `json:"id"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Subtitle  *string   `json:"subtitle"`
	Body      string    `json:"body"`
	Tags      []string  `json:"tags"`
	ImgUrl    *string   `json:"img_url"`
	Published bool      `json:"published"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
