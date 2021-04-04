package repository

import (
	"api-dev-house/src/models"
	"database/sql"
	"fmt"
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

//GetUserById ... retorna um usuário com base no id
func (u Users) GetUserById(id int64) (models.User, error) {
	query, err := u.db.Query("SELECT user_id, name, login, email, created_at FROM tb_users WHERE user_id = ?", id)
	if err != nil {
		return models.User{}, err
	}
	defer query.Close()

	if query.Next() {
		var user models.User
		err = query.Scan(&user.Id, &user.Name, &user.Login, &user.Email, &user.CreatedAt)
		if err != nil {
			return models.User{}, err
		}
		return user, nil
	}

	return models.User{}, nil

}

//UpdateUser ... atualiza dados de um usuário
func (u Users) UpdateUser(id int64, user models.User) error {
	stm, err := u.db.Prepare("UPDATE tb_users set name = ?, email= ?, login = ? WHERE user_id = ?")
	if err != nil {
		return err
	}
	defer stm.Close()

	if _, err := stm.Exec(user.Name, user.Email, user.Login, id); err != nil {
		return err
	}

	return nil

}

//DeleteUser ... remove dados de um usário
func (u Users) DeleteUser(id int64) error {
	stm, err := u.db.Prepare("DELETE FROM tb_users WHERE user_id = ?")
	if err != nil {
		return err
	}
	defer stm.Close()

	if _, err := stm.Exec(id); err != nil {
		return err
	}

	return nil
}

//SearchByLoginOrName ... retorna todos os usuarios que atendem o filtro de nome ou login
func (u Users) SearchByLoginOrName(loginOrName string) ([]models.User, error) {
	loginOrName = fmt.Sprintf("%%%s%%", loginOrName)
	query, err := u.db.Query("SELECT user_id, name, email, login FROM tb_users WHERE login LIKE ? OR name LIKE ? ORDER BY name, login asc", loginOrName, loginOrName)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	var users []models.User
	for query.Next() {
		var user models.User
		if err := query.Scan(&user.Id, &user.Name, &user.Email, &user.Login); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil

}

func (u *Users) SearchByEmail(user models.User) (models.User, error) {
	query, err := u.db.Query("SELECT user_id, password FROM tb_users WHERE email = ?", user.Email)
	if err != nil {
		return models.User{}, err
	}

	var userData models.User
	if query.Next() {
		if err := query.Scan(&userData.Id, &userData.Password); err != nil {
			return models.User{}, err
		}
	}

	return userData, nil
}

//ExitsEmail ... verifica se existe um e-mail cadastrado
func (u Users) ExistEmail(user models.User, isCreatedUser bool) (bool, error) {
	var (
		query *sql.Rows
		err   error
	)

	if isCreatedUser {
		query, err = u.db.Query("SELECT user_id from tb_users where email = ?", user.Email)
	} else {
		query, err = u.db.Query("SELECT user_id from tb_users where email = ? AND user_id != ?", user.Email, user.Id)
	}
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
func (u Users) ExistLogin(user models.User, isCreatedUser bool) (bool, error) {
	var (
		query *sql.Rows
		err   error
	)

	if isCreatedUser {
		query, err = u.db.Query("SELECT user_id FROM tb_users WHERE login = ?", user.Login)
	} else {
		query, err = u.db.Query("SELECT user_id from tb_users where login = ? AND user_id != ?", user.Login, user.Id)
	}

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

//GetPasswordByUserID ... retorna a senha de um usuario pelo ID
func (u *Users) GetPasswordByUserID(userID int64) (string, error) {
	query, err := u.db.Query("SELECT password FROM tb_users WHERE user_id = ?", userID)
	if err != nil {
		return "", err
	}
	defer query.Close()

	var user models.User
	if query.Next() {
		if err := query.Scan(&user.Password); err != nil {
			return "", nil
		}
	}
	return user.Password, nil
}

//UpdatePassword ... atualiza senha do usuario
func (u *Users) UpdatePassword(userID int64, password string) error {
	stm, err := u.db.Prepare("UPDATE tb_users SET password = ? WHERE user_id = ?")
	if err != nil {
		return err
	}
	defer stm.Close()

	if _, err := stm.Exec(password, userID); err != nil {
		return err
	}

	return nil

}
