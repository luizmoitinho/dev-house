package controllers

import (
	"api-dev-house/src/authentication"
	"api-dev-house/src/database"
	"api-dev-house/src/repository"
	"api-dev-house/src/responses"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Follow ... permite que um usuário siga outro
func Follow(w http.ResponseWriter, r *http.Request) {
	followingID, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	userID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadGateway, err)
		return
	}

	if userID == followingID {
		responses.Error(w, http.StatusForbidden, errors.New("não é possivel seguir você mesmo"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryFollow(db)
	isFollow, err := repository.IsFollow(userID, followingID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}

	if isFollow {
		responses.Error(w, http.StatusAccepted, errors.New("usário já está sendo seguido"))
		return
	}

	if err := repository.Follow(userID, followingID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
	return

}

//UnFollow  ... deixa de seguir um usuário
func UnFollow(w http.ResponseWriter, r *http.Request) {
	followingID, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	userID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if userID == followingID {
		responses.Error(w, http.StatusForbidden, errors.New("não é possivel deixar de seguir você mesmo"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryFollow(db)

	if err := repository.UnFollow(userID, followingID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}
