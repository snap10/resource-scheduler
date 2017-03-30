package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

//SetUserRoutes defines the public routes
func SetWebRoutes(router *mux.Router) *mux.Router {
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("public/"))))
	return router
}
