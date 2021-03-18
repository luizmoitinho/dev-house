package repository

import (
	"api-dev-house/src/models"
	"database/sql"
)

//Users ...
type Users struct {
	db *sql.DB
}

//NewRepositoryUser ... instancia uma nova conexao DB de users
func NewRepositoryUser(db *sql.DB) *Users {
	return &Users{db}
}

//Insert ... persiste um novo usuario
func (u Users) Insert(user models.User) (int64, error) {
	stm, err := u.db.Prepare("INSERT INTO tb_users (name, login, email, password) 	VALUES (?,?,?,?)")

	if err != nil {
		return 0, err
	}
	defer stm.Close()

	res, err := stm.Exec(user.Name, user.Login, user.Email, user.Password)

	if err != nil {
		return 0, err
	}

	lastInsertId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertId, nil
}
