package router

import "github.com/gorilla/mux"

type RouteModule interface {
	RegisterRoutes(r *mux.Router)
}
