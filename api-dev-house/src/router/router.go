package router

import (
	"api-dev-house/src/router/routes"

	"github.com/gorilla/mux"
)

//GenerateRouter ... retorna um router com rotas configuradas
func GenerateRouter() *mux.Router {
	router := mux.NewRouter()
	return routes.Config(router)
}
