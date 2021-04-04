package routes

import (
	"api-dev-house/src/controllers"
	"net/http"
)

var followRoutes = []Route{
	{
		URI:           "/users/{id}/follow",
		Method:        http.MethodPost,
		MethodRequest: controllers.Follow,
		Authorization: true,
	},
	{
		URI:           "/users/{id}/unfollow",
		Method:        http.MethodPost,
		MethodRequest: controllers.UnFollow,
		Authorization: true,
	},
}
