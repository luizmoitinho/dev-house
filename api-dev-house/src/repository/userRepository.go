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

//ExitsEmail ... verifica se existe um e-mail cadastrado
func (u Users) ExistEmail(user models.User) (bool, error) {
	query, err := u.db.Query("SELECT user_id from tb_users where email = ?", user.Email)
	if err != nil {
		return false, err
	}
	defer query.Close()

	var id int

	if query.Next() {
		if err := query.Scan(&id); err != nil {
			return false, err
		}
	}

	if id == 0 {
		return false, err
	}
	return true, err

}

//ExitsEmail ... verifica se existe um e-mail cadastrado
func (u Users) ExistLogin(user models.User) (bool, error) {
	query, err := u.db.Query("SELECT user_id FROM tb_users WHERE login = ?", user.Login)
	if err != nil {
		return false, err
	}
	defer query.Close()

	var id int
	if query.Next() {
		if err := query.Scan(&id); err != nil {
			return false, nil
		}
	}

	if id == 0 {
		return false, nil
	}
	return true, nil

}
