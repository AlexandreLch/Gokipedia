package models

import "time"

type Section struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Paragraph string    `json:"paragraph"`
	Position  uint64    `json:"position"`
	Media     string    `json:"media"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
	ParentID  uint64    `json:"parent_id"`
	ArticleID uint64    `json:"article_id"`
}
