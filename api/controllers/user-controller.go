package controllers

import (
	"net/http"

	"github.com/Saqlainabbasi/api/dto"
	"github.com/Saqlainabbasi/api/services"
	"github.com/Saqlainabbasi/api/utils"
)

var (
	UserService services.UserService = services.NewUserService()
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// passing data to service layer for processing
	// services.CreateUser(w, r)

	//user of type dto struct
	user := &dto.User{}
	//passing the data to the utils func
	//this will convert the json data to the specific format like user Struct
	utils.PraseBody(r, user)
	//passing the data to the service layer.....
	newUser, err := UserService.CreateUser(user)

	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, err)
	}
	//sending response back.....
	utils.WriteJson(w, http.StatusOK, newUser)
}

// func to get all the users...
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := UserService.GetUsers()

	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, err)
	}
	//sending response back.....
	utils.WriteJson(w, http.StatusOK, users)
}
