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

//Insert ... insere um novo post
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

//GetByID ... busca um post por id
func (p Posts) GetByID(postID int64) (models.Post, error) {
	query, err := p.db.Query(`
		SELECT p.post_id, p.title, p.content, p.author_id, p.likes, p.created_at, u.login 
			FROM tb_posts p 
			INNER JOIN tb_users u 
				ON u.user_id = p.author_id
		WHERE p.post_id = ?
	`, postID)
	if err != nil {
		return models.Post{}, err
	}
	defer query.Close()

	var post models.Post
	if query.Next() {
		if err := query.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorUser); err != nil {
			return models.Post{}, err
		}
	}
	return post, nil

}

//GetPosts ... busca os seus pots e das pessaos que a pessoa segue
func (p Posts) GetPosts(userID int64) ([]models.Post, error){
	query, err := p.db.Query(`SELECT p.*, u.login 
														FROM tb_posts as p
															INNER JOIN tb_users as u
																ON p.author_id = u.user_id
															INNER JOIN tb_followers as f
																ON f.user_id = p.author_id OR f.following_id = p.author_id
														 WHERE u.user_id = ? OR u.user_id = f.user_id;`, userID)
	
	if err != nil{
		return  nil, err
	}

	var posts []models.Post
	for query.Next(){
		var post models.Post
		if err := query.Scan(&post.Id, &post.Title, &post.Content, &post.AuthorID, &post.Likes, &post.CreatedAt, &post.AuthorUser); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}
