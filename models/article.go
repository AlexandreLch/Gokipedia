package models

import "time"

type Article struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Header    string    `json:"header"`
	Authors   string    `json:"authors"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
}
