package controllers

import (
	"api-dev-house/src/database"
	"api-dev-house/src/models"
	"api-dev-house/src/repository"
	"api-dev-house/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//CreateUser ... cadastrar um novo usuario
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	bodyRequest, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}
	defer db.Close()

	repository := repository.NewRepositoryUser(db)
	user.Id, err = repository.Insert(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}

	responses.JSON(w, http.StatusCreated, user)

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
