package routes

import (
	"net/http"

	"api-dev-house/src/controllers"
)

var userRoutes = []Route{
	{
		URI:           "/users",
		Method:        http.MethodPost,
		MethodRequest: controllers.CreateUser,
		Authorization: false,
	},
	{
		URI:           "/users",
		Method:        http.MethodGet,
		MethodRequest: controllers.GetUsers,
		Authorization: false,
	},
	{
		URI:           "/user/{id}",
		Method:        http.MethodGet,
		MethodRequest: controllers.GetUser,
		Authorization: false,
	},
	{
		URI:           "/users/{id}",
		Method:        http.MethodPut,
		MethodRequest: controllers.UpdateUser,
		Authorization: false,
	},
	{
		URI:           "/users/{id}",
		Method:        http.MethodDelete,
		MethodRequest: controllers.DeleteUser,
		Authorization: false,
	},
}
