package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leangeder/chatops2/pkg/server/handler"
	"github.com/leangeder/chatops2/pkg/server/logger"
)

// Route type description
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes contains all routes
type Routes []Route

var routes = Routes{
	Route{
		"GetPeople",
		"GET",
		"/people",
		handler.GetTest,
	},
	Route{
		"GetPerson",
		"GET",
		"/people/{id}",
		handler.GetTest,
	},
	Route{
		"GetLogs",
		"GET",
		"/k8s/logs/{name}",
		handler.DeploymentToPreview,
	},
}

// NewRouter builds and returns a new router from routes
func NewRouterV1() *mux.Router {
	// When StrictSlash == true, if the route path is "/path/", accessing "/path" will perform a redirect to the former and vice versa.
	router := mux.NewRouter().StrictSlash(true)
	router.Use(logger.Logger)
	sub := router.PathPrefix("/v1").Subrouter()

	for _, route := range routes {
		sub.
			HandleFunc(route.Pattern, route.HandlerFunc).
			Name(route.Name).
			Methods(route.Method)
	}

	return router
}
