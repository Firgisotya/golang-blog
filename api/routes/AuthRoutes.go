package routes

import (
	"go-blog/api/handlers"
	"net/http"
)

var authRoute = []Route{
	Route{
		URI:     "/auth/login",
		Method:  http.MethodPost,
		Handler: handlers.LoginHandler,
		Auth:    false,
	},
}