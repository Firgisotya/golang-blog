package routes

import (
	"go-blog/api/middlewares"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Route struct {
	URI     string
	Method  string
	Handler func(w http.ResponseWriter, r *http.Request)
	Auth    bool
}

func Load() []Route {
	routes := welcomeRoute
	routes = append(routes, authRoute...)

	return routes
}

func SetupMiddleware(r *mux.Router) *mux.Router {
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	for _, route := range Load() {
		api := r.PathPrefix("/api").Subrouter()

		if route.Auth {
			api.HandleFunc(route.URI,
				middlewares.SetMiddlewareCors(
					middlewares.SetMiddlewareAuthentication(
						route.Handler,
					),
				),
			).Methods(route.Method)
		} else {
			api.HandleFunc(route.URI,
				middlewares.SetMiddlewareCors(
					route.Handler,
				),
			).Methods(route.Method)
		}
	}

	return r

}
