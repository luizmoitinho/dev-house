package repository

import (
	"api-dev-house/src/models"
	"database/sql"
)

//Posts ... repository de publicações
type Posts struct {
	db *sql.DB
}

func NewRepositoryPosts(db *sql.DB) *Posts {
	return &Posts{db}
}

func (p Posts) Insert(post models.Post) (int64, error) {
	stm, err := p.db.Prepare("INSERT INTO tb_posts (title, content, author_id) VALUES (?,?,?)")
	if err != nil {
		return 0, nil
	}
	defer stm.Close()

	result, err := stm.Exec(post.Title, post.Content, post.AuthorID)
	if err != nil {
		return 0, nil
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return int64(lastId), nil

}
