package controllers

import (
	"net/http"

	"github.com/Saqlainabbasi/api/config"
	"github.com/Saqlainabbasi/api/models"
	"github.com/Saqlainabbasi/api/utils"
	"gorm.io/gorm"
)

// declare the db variable of type gorm.DB
var db *gorm.DB

// declare models
var User *models.User

func CreateUser(w http.ResponseWriter, r *http.Request) {
	//get DB
	db = config.GetDB()
	//
	user := &models.User{}
	utils.PraseBody(r, user)
	//query to add data to the user
	resp := db.Create(user)
	if resp.Error != nil {
		utils.WriteJson(w, http.StatusInternalServerError, resp.Error)
	}
	//send the responce back
	utils.WriteJson(w, http.StatusOK, user)
}

// func to get all the users...
func GetUsers(w http.ResponseWriter, r *http.Request) {
	db = config.GetDB()
	//query to get all the users
	var users []models.User
	db.Find(&users)
	utils.WriteJson(w, http.StatusOK, users)
}
