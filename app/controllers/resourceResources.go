package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/snap10/resource-scheduler/app/data"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/snap10/resource-scheduler/app/common"
	"github.com/snap10/resource-scheduler/app/models"

	"log"
)

type (
	ResourceListData struct {
		Data []models.Resource `json:"data"`
	}
	ReservationData struct {
		Data models.Reservation `json:"data"`
	}
	ResourceData struct {
		Data models.Resource `json:"data"`
	}
	UserIdData struct {
		Data models.User `json:"data"`
	}
)

func getStringFromContext(r *http.Request, key interface{}) string {
	if val := context.Get(r, key); val != nil {
		return val.(string)
	}
	return "no value"
}

//Register is Handler for HTTP Post "/user/register"

func ListResources(w http.ResponseWriter, r *http.Request) {
	log.Printf("[controller.ListResources]: %s\n", "New request from "+r.RemoteAddr+"by"+getStringFromContext(r, "user"))
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("resources")
	resourceCollection := &data.ResourceCollection{C: c}
	resources, err := resourceCollection.FindAllResources()
	if err != nil {
		common.DisplayAppError(w, err, "Error while getting resources", http.StatusInternalServerError)
		return
	}
	j, err := json.Marshal(ResourceListData{Data: *resources})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func CreateResource(w http.ResponseWriter, r *http.Request) {
	log.Printf("[controller.CreateResource]: %s\n", "New request from "+r.RemoteAddr)
	var dataResource ResourceData
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Resource data", 500)
		return

	}
	resource := &dataResource.Data

	context := NewContext()
	defer context.Close()
	c := context.DbCollection("resources")
	resourceCollection := &data.ResourceCollection{C: c}
	err = resourceCollection.CreateResource(resource)
	if err != nil {
		common.DisplayAppError(w, err, "Error while getting resources", http.StatusInternalServerError)
		return
	}
	j, err := json.Marshal(ResourceData{Data: *resource})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}
func GetResource(w http.ResponseWriter, r *http.Request) {
	common.DisplayAppError(w, nil, "Not implemented yet!", http.StatusNotImplemented)
}
func UpdateResource(w http.ResponseWriter, r *http.Request) {
	common.DisplayAppError(w, nil, "Not implemented yet!", http.StatusNotImplemented)
}
func DeleteResource(w http.ResponseWriter, r *http.Request) {
	common.DisplayAppError(w, nil, "Not implemented yet!", http.StatusNotImplemented)
}
func ListResourceUsers(w http.ResponseWriter, r *http.Request) {
	common.DisplayAppError(w, nil, "Not implemented yet!", http.StatusNotImplemented)
}
func AssignUserToResource(w http.ResponseWriter, r *http.Request) {
	log.Printf("[controller.CreateResource]: %s\n", "New request from "+r.RemoteAddr)
	vars := mux.Vars(r)
	rid := vars["rid"]

	resource := &models.Resource{Id: bson.ObjectIdHex(rid)}
	var userIdData UserResource
	err := json.NewDecoder(r.Body).Decode(&userIdData)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Userid data", 500)
		return

	}
	userID := &userIdData.Data.Id
	log.Printf("User id %s:", userID)
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("resources")
	resourceCollection := &data.ResourceCollection{C: c}
	updatedResource, err := resourceCollection.UpdateResource(resource, "users", userID, "$addToSet")
	if err != nil {
		common.DisplayAppError(w, err, "Error while getting resources", http.StatusInternalServerError)
		return
	}
	j, err := json.Marshal(ResourceData{Data: *updatedResource})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}
func UpdateUserOnResource(w http.ResponseWriter, r *http.Request) {
	common.DisplayAppError(w, nil, "Not implemented yet!", http.StatusNotImplemented)
}
func RemoveUserFromResource(w http.ResponseWriter, r *http.Request) {
	common.DisplayAppError(w, nil, "Not implemented yet!", http.StatusNotImplemented)
}
func ListReservations(w http.ResponseWriter, r *http.Request) {
	common.DisplayAppError(w, nil, "Not implemented yet!", http.StatusNotImplemented)
}
func CreateReservation(w http.ResponseWriter, r *http.Request) {
	log.Printf("[controller.CreateReservation]: %s\n", "New request from "+r.RemoteAddr)
	var dataReservation ReservationData
	err := json.NewDecoder(r.Body).Decode(&dataReservation)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Resource data", 500)
		return

	}
	reservation := &dataReservation.Data

	context := NewContext()
	defer context.Close()
	c := context.DbCollection("reservations")
	reservationCollection := &data.ReservationCollection{C: c}
	err = reservationCollection.CreateReservation(reservation)
	if err != nil {
		common.DisplayAppError(w, err, "Error while getting resources", http.StatusInternalServerError)
		return
	}
	createdReservation, err := json.Marshal(ReservationData{Data: *reservation})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(createdReservation)
}
func GetReservation(w http.ResponseWriter, r *http.Request) {
	common.DisplayAppError(w, nil, "Not implemented yet!", http.StatusNotImplemented)
}
func UpdateReservation(w http.ResponseWriter, r *http.Request) {
	common.DisplayAppError(w, nil, "Not implemented yet!", http.StatusNotImplemented)
}
func DeleteReservation(w http.ResponseWriter, r *http.Request) {
	log.Printf("[controller.DeleteReservation]: %s\n", "New request from "+r.RemoteAddr)
	vars := mux.Vars(r)
	reservationid := vars["reservationid"]
	reservation := &models.Reservation{Id: bson.ObjectIdHex(reservationid)}
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("reservations")
	reservationCollection := &data.ReservationCollection{C: c}
	err := reservationCollection.DeleteReservation(reservation)
	if err != nil {
		common.DisplayAppError(w, err, "Error while getting resources", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	w.Write(nil)
}
