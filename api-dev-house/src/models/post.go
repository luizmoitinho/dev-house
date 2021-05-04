package models

import (
	"errors"
	"strings"
	"time"
)

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

//Prepare ... valida e formata dados do post.s
func (p *Post) Prepare() error {
	if err := p.validate(); err != nil {
		return err
	}

	p.format()
	return nil
}

func (p *Post) validate() error {
	if p.Title == "" {
		return errors.New("O titulo é obrigatório.")
	}
	if p.Content == "" {
		return errors.New("O contéudo é obrigatório.")
	}

	return nil

}

func (p *Post) format() {
	p.Title = strings.TrimSpace(p.Title)
	p.Content = strings.TrimSpace(p.Content)

}
