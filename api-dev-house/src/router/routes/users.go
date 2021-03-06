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
		Authorization: true,
	},
	{
		URI:           "/user/{id}",
		Method:        http.MethodGet,
		MethodRequest: controllers.GetUser,
		Authorization: true,
	},
	{
		URI:           "/users/{id}",
		Method:        http.MethodPut,
		MethodRequest: controllers.UpdateUser,
		Authorization: true,
	},
	{
		URI:           "/users/{id}",
		Method:        http.MethodDelete,
		MethodRequest: controllers.DeleteUser,
		Authorization: true,
	},
	{
		URI:           "/users/{id}/password",
		Method:        http.MethodPost,
		MethodRequest: controllers.UpdatePassword,
		Authorization: true,
	},
}
