package data

import (
	"github.com/snap10/resource-scheduler/app/models"
	"golang.org/x/crypto/bcrypt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserCollection struct {
	C *mgo.Collection
}

func (ucoll *UserCollection) CreateUser(user *models.User) error {
	obj_id := bson.NewObjectId()
	user.Id = obj_id
	hpass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.HashPassword = hpass
	//clear the incoming cleartext pw
	user.Password = ""
	err = ucoll.C.Insert(&user)
	return err
}

func (ucoll *UserCollection) Login(user models.User) (u models.User, err error) {
	err = ucoll.C.Find(bson.M{"email": user.Email}).One(&u)
	if err != nil {
		return
	}
	//validate pw
	err = bcrypt.CompareHashAndPassword(u.HashPassword, []byte(user.Password))
	if err != nil {
		u = models.User{}
	}
	return
}
