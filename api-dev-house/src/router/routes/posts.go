package routes

import (
	"api-dev-house/src/controllers"
	"net/http"
)

var postsRoutes = []Route{
	{
		URI:           "/posts",
		Method:        http.MethodPost,
		MethodRequest: controllers.CreatePost,
		Authorization: true,
	},
	{
		URI:           "/posts",
		Method:        http.MethodGet,
		MethodRequest: controllers.GetPosts,
		Authorization: true,
	},
	{
		URI:           "/posts/{id}",
		Method:        http.MethodGet,
		MethodRequest: controllers.GetPost,
		Authorization: true,
	},
	{
		URI:           "/posts/{id}",
		Method:        http.MethodPut,
		MethodRequest: controllers.UpdatePost,
		Authorization: true,
	},
	{
		URI:           "/posts/{id}",
		Method:        http.MethodDelete,
		MethodRequest: controllers.DeletePost,
		Authorization: true,
	},
}
