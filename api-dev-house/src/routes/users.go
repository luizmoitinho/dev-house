package routes

import "net/http"

var userRoutes = []Route{
	{
		URI: "/users",
		Method: http.MethodPost,
		MethodRequest: func(w http.ResponseWriter, r * http.Request){
			
		},
		Authorization: false,
	},
	{
		URI: "/users",
		Method: http.MethodGet,
		MethodRequest: func(w http.ResponseWriter, r * http.Request){
			
		},
		Authorization: false,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodGet,
		MethodRequest: func(w http.ResponseWriter, r * http.Request){
			
		},
		Authorization: false,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodGet,
		MethodRequest: func(w http.ResponseWriter, r * http.Request){
			
		},
		Authorization: false,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodPut,
		MethodRequest: func(w http.ResponseWriter, r * http.Request){
			
		},
		Authorization: false,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodDelete,
		MethodRequest: func(w http.ResponseWriter, r * http.Request){
			
		},
		Authorization: false,
	}
}


