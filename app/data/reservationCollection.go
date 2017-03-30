package data

import (
	"log"

	"github.com/snap10/resource-scheduler/app/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ReservationCollection struct {
	C *mgo.Collection
}

func (coll *ReservationCollection) FindAllReservations() (reservations *[]models.Reservation, err error) {

	reservationsSlice := make([]models.Reservation, 0, 10)
	err = coll.C.Find(bson.M{}).All(&reservationsSlice)
	if err != nil {
		return nil, err
	}
	return &reservationsSlice, nil
}

func (coll *ReservationCollection) CreateReservation(reservation *models.Reservation) error {
	objID := bson.NewObjectId()
	reservation.Id = objID

	err := coll.C.Insert(&reservation)
	return err
}

func (coll *ReservationCollection) UpdateReservation(reservation *models.Reservation, key string, data interface{}, mgoOperator string) (updatedReservation *models.Reservation, err error) {
	log.Printf("data.[UpdateReservation]- key %s, data %s", key, data)
	change := mgo.Change{
		Update:    bson.M{mgoOperator: bson.M{key: data}},
		ReturnNew: true,
	}
	info, err := coll.C.FindId(reservation.Id).Apply(change, &updatedReservation)
	log.Print("info: %s", info.Updated)
	return reservation, err
}

func (coll *ReservationCollection) DeleteReservation(reservation *models.Reservation) (err error) {
	log.Printf("data.[DeleteReservation]- with id %s", reservation.Id)

	err = coll.C.RemoveId(reservation.Id)
	return err
}
