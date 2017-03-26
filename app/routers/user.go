package routers

import (
	"github.com/gorilla/mux"
	"github.com/snap10/resource-scheduler/app/controllers"
)

//SetUserRoutes defines the public routes
func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/user/register", controllers.Register).Methods("POST")
	router.HandleFunc("/user/login", controllers.Login).Methods("POST")
	return router
}
