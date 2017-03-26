package controllers

import "github.com/snap10/resource-scheduler/app/models"

type (
	//UserResource For Post - /user/register
	UserResource struct {
		Data models.User `json:"data"`
	}

	//LoginResource For Post -/user/login
	LoginResource struct {
		Data LoginModel `json:"data"`
	}

	//AuthUserResource is Response resource for /user/login
	AuthUserResource struct{
		Data AuthUserModel `Json:"data"`
	}

	//LoginModel is for Authentication against /user/login
	LoginModel struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}

	//AuthUserModel is for Access with Token
	AuthUserModel struct {
		User models.User `json:"user"`
		Token string `json:"token"`
	}
)
