package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

//Logger ... informa no terminal informações da requisicao
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

//Authenticate ... verifica se o usuario que esta fazendo a requisição está autenticado
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Autenticando ...")
		next(w, r)
	}
}
