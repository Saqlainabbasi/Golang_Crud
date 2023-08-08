package services

import (
	"net/http"

	"github.com/Saqlainabbasi/api/models"
	"github.com/Saqlainabbasi/api/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	//mapping the json data to the DB model format....
	// variable of type Struct
	user := &models.User{}
	// use the utils function formate the data
	utils.PraseBody(r, user)
	//pass data to database layer to store in db
	resp := user.CreateUser()
	if resp.Error != nil {
		utils.WriteJson(w, http.StatusInternalServerError, resp.Error)
	}
	// send the responce back
	utils.WriteJson(w, http.StatusOK, user)
}

// service function to get all the users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	//pass data to database layer to store in db
	resp, data := models.GetUsers()
	if resp.Error != nil {
		utils.WriteJson(w, http.StatusInternalServerError, resp.Error)
	}
	// send the responce back
	utils.WriteJson(w, http.StatusOK, data)
}
