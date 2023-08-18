package controllers

// import (
// 	"net/http"

// 	"github.com/Saqlainabbasi/api/services"
// )

// func CreateUser(w http.ResponseWriter, r *http.Request) {
// 	// passing data to service layer for processing
// 	services.CreateUser(w, r)
// }

// // func to get all the users...
// func GetUsers(w http.ResponseWriter, r *http.Request) {
// 	services.GetUsers(w, r)
// }

// ***********Old Function ************
// func CreateUser(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("a")
// 	//get DB
// 	db = config.GetDB()
// 	//
// 	user := &models.User{}
// 	utils.PraseBody(r, user)
// 	//query to add data to the user
// 	resp := db.Create(user)
// 	if resp.Error != nil {
// 		utils.WriteJson(w, http.StatusInternalServerError, resp.Error)
// 	}
// 	//send the responce back
// 	utils.WriteJson(w, http.StatusOK, user)
// }

// func CreateUser(w http.ResponseWriter, r *http.Request) {
// fmt.Println("check")
// user := &models.User{}
// utils.PraseBody(r, user)
// // calling the model method to create user
// resp := user.CreateUser()
// if resp.Error != nil {
// 	utils.WriteJson(w, http.StatusInternalServerError, resp.Error)
// }
// //send the responce back
// utils.WriteJson(w, http.StatusOK, user)
// }
