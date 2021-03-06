package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/snap10/resource-scheduler/app/common"
	"github.com/snap10/resource-scheduler/app/controllers"
)

//SetUserRoutes defines the public routes
func SetResourceRoutes(router *mux.Router) *mux.Router {
	resourceRouter := mux.NewRouter()
	resourceRouter.HandleFunc("/resources", controllers.ListResources).Methods("GET")
	resourceRouter.HandleFunc("/resources", controllers.CreateResource).Methods("POST")
	resourceRouter.HandleFunc("/resources/{rid}", controllers.GetResource).Methods("GET")
	resourceRouter.HandleFunc("/resources/{rid}", controllers.UpdateResource).Methods("PUT")
	resourceRouter.HandleFunc("/resources/{rid}", controllers.DeleteResource).Methods("DELETE")
	resourceRouter.HandleFunc("/resources/{rid}/users", controllers.ListResourceUsers).Methods("GET")
	resourceRouter.HandleFunc("/resources/{rid}/users", controllers.AssignUserToResource).Methods("POST")
	resourceRouter.HandleFunc("/resources/{rid}/users/{uid}", controllers.UpdateUserOnResource).Methods("PUT")
	resourceRouter.HandleFunc("/resources/{rid}/users/{uid}", controllers.RemoveUserFromResource).Methods("DELETE")
	resourceRouter.HandleFunc("/resources/{rid}/reservations", controllers.ListReservations).Methods("GET")
	resourceRouter.HandleFunc("/resources/{rid}/reservations", controllers.CreateReservation).Methods("POST")
	resourceRouter.HandleFunc("/resources/{rid}/reservations/{reservid}", controllers.GetReservation).Methods("GET")
	resourceRouter.HandleFunc("/resources/{rid}/reservations/{reservid}", controllers.UpdateReservation).Methods("PUT")
	resourceRouter.HandleFunc("/resources/{rid}/reservations/{reservid}", controllers.DeleteReservation).Methods("DELETE")
	router.PathPrefix("/resources").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(resourceRouter),
	))
	return router
}
