package routes

import (
	"go-blog/api/handlers"
	"net/http"
)

var welcomeRoute = []Route{
	Route{
		URI:     "/",
		Method:  http.MethodGet,
		Handler: handlers.Welcome,
		Auth:    false,
	},
}