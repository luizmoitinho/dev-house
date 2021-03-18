package database

import (
	"api-dev-house/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver
)

//Connect ... abre e retorna conexao com banco de dados
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.ConnectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil

}
