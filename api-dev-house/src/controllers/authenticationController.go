package controllers

import (
	"api-dev-house/src/authentication"
	"api-dev-house/src/database"
	"api-dev-house/src/models"
	"api-dev-house/src/repository"
	"api-dev-house/src/responses"
	"api-dev-house/src/security"
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
	userDB, err := repository.SearchByEmail(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err := security.ComparePasswords(userDB.Password, user.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, _ := authentication.GenerateToken(userDB.Id)
	w.Write([]byte(token))

}
