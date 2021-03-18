package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Route .. representa a estutura das rotas da api
type Route struct {
	URI           string
	Method        string
	MethodRequest func(http.ResponseWriter, *http.Request)
	Authorization bool
}

//Config ... configura todas as rotas dentro do router
func Config(r *mux.Router) *mux.Router {
	routes := userRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.MethodRequest).Methods(route.Method)
	}

	return r

}
