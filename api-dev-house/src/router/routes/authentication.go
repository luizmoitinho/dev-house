package routes

import (
	"api-dev-house/src/controllers"
	"net/http"
)

var authenticationRoutes = Route{

	URI:           "/login",
	Method:        http.MethodPost,
	MethodRequest: controllers.Login,
	Authorization: false,
}
