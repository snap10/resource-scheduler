package models

import "gopkg.in/mgo.v2/bson"

type (
	User struct {
		Id           bson.ObjectId `bson:"_id,omitempty" json:"id"`
		FirstName    string        `json:"firstname"`
		LastName     string        `json:"lastname"`
		Email        string        `json:"email"`
		Password     string        `json:"password,omitempty"`
		HashPassword []byte        `json:"hashpassword,omitempty"`
	}
	Resource struct {
		Id    bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name  string        `json:"name"`
		Users []User        `json:"users"`
	}
	Reservation struct {
		Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		UserID      bson.ObjectId `bson:"userid" json:"userid"`
		ResourceID  bson.ObjectId `bson:"resourceid" json:"resourceid"`
		Name        string        `json:"name"`
		Description string        `json:"description"`
		Date        int64         `json:"date"`
		Duration    int32         `json:"duration"`
		AllDay      bool          `json:"all_day"`
	}
)
