package controller

import "net/http"

//CreateUser ... cadastrar um novo usuario
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário!"))
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
