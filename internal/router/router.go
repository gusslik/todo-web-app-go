package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Setup(routeModules []RouteModule) http.Handler {
	r := mux.NewRouter()

	// Register all routes
	for _, module := range routeModules {
		module.RegisterRoutes(r)
	}

	return r
}
