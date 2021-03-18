package router

import "github.com/gorilla/mux"

//GenerateRouter ... retorna um router com rotas configuradas
func GenerateRouter() *mux.Router {
	return mux.NewRouter()
}
