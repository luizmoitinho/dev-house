package controllers

import (
	"api-dev-house/src/authentication"
	"api-dev-house/src/database"
	"api-dev-house/src/models"
	"api-dev-house/src/repository"
	"api-dev-house/src/responses"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//CreatePost ... adiciona uma nova publicacao
func CreatePost(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err = json.Unmarshal(bodyRequest, &post); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	post.AuthorID = userID
	if err := post.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	post.AuthorID = userID
	repository := repository.NewRepositoryPosts(db)
	post.Id, err = repository.Insert(post)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, post)

}

//GetPosts ... exibe posts no feed do user
func GetPosts(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryPosts(db)

	posts, err := repository.GetPosts(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	fmt.Println(posts)

	responses.JSON(w, http.StatusOK, posts)

}

//GetPost ... exibe uma publicação espeficica
func GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseInt(params["id"], 10, 64)
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

	repository := repository.NewRepositoryPosts(db)
	post, err := repository.GetByID(postID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, post)

}

//UpdatePost ... atualiza uma publicação
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	postID, err := strconv.ParseInt(params["id"], 10, 64)
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

	repository := repository.NewRepositoryPosts(db)
	post, err := repository.GetByID(postID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if userID != post.AuthorID {
		responses.Error(w, http.StatusForbidden, errors.New("não é possivel atualizar publicações de terceiros"))
		return
	}

	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = json.Unmarshal(bodyRequest, &post); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := post.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := repository.Update(postID, post); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

//DeletePost ... deleta uma publicação
func DeletePost(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	postID, err := strconv.ParseInt(params["id"], 10, 64)
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

	repository := repository.NewRepositoryPosts(db)
	post, err := repository.GetByID(postID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if userID != post.AuthorID {
		responses.Error(w, http.StatusForbidden, errors.New("não é possivel remover publicações de terceiros"))
		return
	}

	if err := repository.Delete(postID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}
