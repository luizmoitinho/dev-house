package routes

import (
	"api-dev-house/src/controllers"
	"net/http"
)

var followRoutes = []Route{
	{
		URI:           "/follow/{id}",
		Method:        http.MethodPost,
		MethodRequest: controllers.Follow,
		Authorization: true,
	},
	{
		URI:           "/unfollow/{id}",
		Method:        http.MethodPost,
		MethodRequest: controllers.UnFollow,
		Authorization: true,
	},
	{
		URI:           "/followers/{id}",
		Method:        http.MethodGet,
		MethodRequest: controllers.GetFollowers,
		Authorization: true,
	},
}
