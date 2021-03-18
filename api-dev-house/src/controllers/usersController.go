package controllers

import (
	"api-dev-house/src/database"
	"api-dev-house/src/models"
	"api-dev-house/src/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//CreateUser ... cadastrar um novo usuario
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repository := repository.NewRepositoryUser(db)
	Id, err := repository.Insert(user)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(fmt.Sprintf("id inserido: %d", Id)))
}

//GetUsers ... retorna todos os usuarios
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuários!"))
}

//GetUser ... retorna um usuario
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar um usuário!"))
}

//UpdateUser ... atualiza dados de um usuario
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizar usuário!"))
}

//UpdateUser ... remove um usuario
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Remover usuário!"))
}
