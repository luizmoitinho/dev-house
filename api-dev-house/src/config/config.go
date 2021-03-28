package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//String de conexão do banco mysql
	ConnectionString = ""

	//Porta ondeaa api estará atuando
	Port = 0

	//Secret key é a chave para assinar o jwt
	SecretKey []byte
)

//Load ... carregar as variáveis de ambiente
func Load() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal("config.Load(): ", err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 9000
	}
	ConnectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))

}
