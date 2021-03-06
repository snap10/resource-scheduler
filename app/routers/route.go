package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	// Routes for the User entity
	router = SetUserRoutes(router)
	// Routes for the Task entity
	router = SetResourceRoutes(router)
	// Routes for weg-app... setup last to prreservation errors√
	router = SetWebRoutes(router)
	return router
}
