package repository

import (
	"api-dev-house/src/models"
	"database/sql"
)

//Follow ...
type Follow struct {
	db *sql.DB
}

//NewRepositoryUser ... instancia uma nova conexao DB de users
func NewRepositoryFollow(db *sql.DB) *Follow {
	return &Follow{db}
}

//Follow ... faz com que um usuário siga outro usuario
func (f *Follow) Follow(userID, followingID int64) error {
	stm, err := f.db.Prepare("INSERT INTO tb_followers (user_id, following_id) VALUES (?,?)")
	if err != nil {
		return err
	}
	defer stm.Close()

	if _, err := stm.Exec(userID, followingID); err != nil {
		return err
	}

	return nil
}

//UnFollow ... para de seguir um determinando usuário
func (f *Follow) UnFollow(userID, followingID int64) error {
	stm, err := f.db.Prepare("DELETE FROM tb_followers WHERE user_id = ? AND following_id = ?")

	if err != nil {
		return nil
	}
	defer stm.Close()

	if _, err := stm.Exec(userID, followingID); err != nil {
		return err
	}
	return nil

}

//IsFollow ... verifica se um usuário já segue um outro usuario
func (f *Follow) IsFollow(userID, followingID int64) (bool, error) {
	query, err := f.db.Query("SELECT true FROM tb_followers WHERE user_id = ? AND following_id = ? ", userID, followingID)
	if err != nil {
		return false, err
	}
	defer query.Close()

	if query.Next() {
		var isFollow bool
		if err := query.Scan(&isFollow); err != nil {
			return false, err
		}
		return isFollow, nil
	}
	return false, nil
}

//GetFollowers ... retornar todos os seguidores de um usuário
func (f *Follow) GetFollowers(userID int64) ([]models.User, error) {

	query, err := f.db.Query(`SELECT u.user_id, u.name, u.login, u.email, u.created_at
															FROM tb_users as u 
															INNER JOIN tb_followers as f
															ON u.user_id = f.following_id
															WHERE f.user_id = ?
														`, userID)

	if err != nil {
		return nil, err
	}
	defer query.Close()

	var followers []models.User
	for query.Next() {
		var follower models.User
		if err := query.Scan(&follower.Id, &follower.Name, &follower.Login, &follower.Email, &follower.CreatedAt); err != nil {
			return nil, err
		}
		followers = append(followers, follower)
	}

	return followers, nil

}
