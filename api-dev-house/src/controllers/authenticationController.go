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

//Login ... realiza a autenticação de um usuario na api
func Login(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	if err := json.Unmarshal(bodyRequest, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}
	defer db.Close()

	repository := repository.NewRepositoryUser(db)
	user.Id, err = repository.Login(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)

}
