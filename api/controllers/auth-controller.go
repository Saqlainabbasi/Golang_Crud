package controllers

import (
	"net/http"

	"github.com/Saqlainabbasi/api/dto"
	"github.com/Saqlainabbasi/api/models"
	"github.com/Saqlainabbasi/api/services"
	"github.com/Saqlainabbasi/api/utils"
)

var (
	AuthService services.AuthService = services.NewAuthService(models.NewDAO(), services.NewTokenManager("mysecret"), services.NewAMQP())
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	// passing data to service layer for processing
	// services.CreateUser(w, r)

	//user of type dto struct
	user := &dto.User{}
	//passing the data to the utils func
	//this will convert the json data to the specific format like user Struct
	utils.PraseBody(r, user)
	//passing the data to the service layer.....
	newUser, err := AuthService.SignUp(user)

	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, err)
	}
	//sending response back.....
	utils.WriteJson(w, http.StatusOK, newUser)
}
func SignIn(w http.ResponseWriter, r *http.Request) {

	body := &dto.User{}
	//passing the data to the utils func
	//this will convert the json data to the specific format like user Struct
	utils.PraseBody(r, body)
	user, err := AuthService.SignIn(body.Email, body.Paswrd)
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, err)
	}
	utils.WriteJson(w, http.StatusOK, user)

}

// func to get all the users...
// func GetUsers(w http.ResponseWriter, r *http.Request) {
// 	users, err := UserService.GetUsers()

// 	if err != nil {
// 		utils.WriteJson(w, http.StatusInternalServerError, err)
// 	}
// 	//sending response back.....
// 	utils.WriteJson(w, http.StatusOK, users)
// }
