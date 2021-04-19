package models

import "time"

//Post ... dados de uma publicação
type Post struct {
	Id         int64     `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorUser string    `json:"author_user,omitempty"`
	AuthorID   int64     `json:"author_id,omitempty"`
	Likes      int64     `json:"likes"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}
