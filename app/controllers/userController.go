package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/snap10/resource-scheduler/app/common"
	"github.com/snap10/resource-scheduler/app/data"
	"github.com/snap10/resource-scheduler/app/models"
)

//Register is Handler for HTTP Post "/user/register"

func Register(w http.ResponseWriter, r *http.Request) {
	var dataResource UserResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid User data", 500)
		return

	}
	user := &dataResource.Data
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	repo := &data.UserCollection{C: c}
	//Insert User document
	repo.CreateUser(user)
	//clean-up the hashpassword to eliminate it from ResponseWriter
	user.HashPassword = nil
	j, err := json.Marshal(UserResource{Data: *user})
	if err != nil {
		common.DisplayAppError(w, err, "Unexpected Error has occured", 500)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)

	}
}

//Login is Handler for HTTP Post "/user/login" with username and Hashpassword

func Login(w http.ResponseWriter, r *http.Request) {
	var dataResource LoginResource
	var token string
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Login data", 500)
		return
	}
	loginModel := &dataResource.Data
	loginUser := models.User{
		Email:    loginModel.Email,
		Password: loginModel.Password,
	}
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	repo := &data.UserCollection{c}
	//Authenticate the login user
	user, err := repo.Login(loginUser)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid login credentials", 401)
		return
	}
	token, err = common.GenerateJWT(user.Email, "member")
	if err != nil {
		common.DisplayAppError(w, err, "Error while creating token", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	user.HashPassword = nil
	authUser := AuthUserModel{
		User:  user,
		Token: token,
	}
	//clean-up the hashpassword to eliminate it from ResponseWriter
	user.HashPassword = nil
	j, err := json.Marshal(AuthUserResource{Data: authUser})
	if err != nil {
		common.DisplayAppError(w, err, "Unexpected Error has occured", 500)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
