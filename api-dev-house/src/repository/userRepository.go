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
func (u Users) Insert(usuario models.User) (int64, error) {

	return 0, nil
}
