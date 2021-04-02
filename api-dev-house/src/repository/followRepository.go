package repository

import (
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

func (f *Follow) isFollow(userID, followingID int64) (bool, error) {
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
