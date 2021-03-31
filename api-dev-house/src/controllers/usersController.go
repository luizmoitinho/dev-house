package controllers

import (
	"api-dev-house/src/authentication"
	"api-dev-house/src/database"
	"api-dev-house/src/models"
	"api-dev-house/src/repository"
	"api-dev-house/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

	if err := user.Prepare(true); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := validateUniqueDataUser(user, true); err != nil {
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
		return
	}

	responses.JSON(w, http.StatusCreated, user)

}

//GetUser ... retorna um usuario
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryUser(db)
	user, err := repository.GetUserById(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	if user.Id == 0 {
		responses.Error(w, http.StatusNotFound, errors.New("Usuarios não encontrado"))
		return
	}
	responses.JSON(w, http.StatusOK, user)

}

//GetUsers ... retorna todos os usuarios
func GetUsers(w http.ResponseWriter, r *http.Request) {
	loginOrName := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryUser(db)

	users, err := repository.SearchByLoginOrName(loginOrName)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)

}

//UpdateUser ... atualiza dados de um usuario
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	userIDToken, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userIDToken != userID {
		responses.Error(w, http.StatusForbidden, errors.New("não é possível manipular usuário de terceiros"))
		return
	}

	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err := json.Unmarshal(bodyRequest, &user); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err := user.Prepare(false); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	user.Id = userID
	if err := validateUniqueDataUser(user, false); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryUser(db)
	if err = repository.UpdateUser(userID, user); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

//UpdateUser ... remove um usuario
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userIDToken, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userIDToken != userID {
		responses.Error(w, http.StatusForbidden, errors.New("não é possível manipular usuário de terceiros"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryUser(db)

	if err := repository.DeleteUser(userID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func validateUniqueDataUser(user models.User, isCreateUser bool) error {

	db, err := database.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	repository := repository.NewRepositoryUser(db)
	if err != nil {
		return err
	}

	loginValid, err := repository.ExistLogin(user, isCreateUser)
	if err != nil {
		return err
	} else if loginValid {
		return errors.New("Login já existe na plataforma")
	}

	emailValid, err := repository.ExistEmail(user, isCreateUser)
	if err != nil {
		return err
	} else if emailValid {
		return errors.New("E-mail já existe na plataforma")
	}

	return nil

}
