package routes

import (
	"api-dev-house/src/middlewares"
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
	routes := append(userRoutes, followRoutes...)

	routes = append(routes, authenticationRoutes)
	for _, route := range routes {
		if route.Authorization {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authenticate(route.MethodRequest)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.MethodRequest)).Methods(route.Method)
		}
	}

	return r

}
