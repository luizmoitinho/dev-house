package routes

import "net/http"

//Route .. representa a estutura das rotas da api
type Route struct {
	URI           string
	Method        string
	MethodRequest func(http.ResponseWriter, *http.Request)
	Authorization bool
}
