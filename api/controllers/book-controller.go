package controllers

import (
	"net/http"

	"github.com/Saqlainabbasi/api/services"
)

// func CreateBook(w http.ResponseWriter, r *http.Request) {
// 	// passing data to service layer for processing
// 	services.CreateUser(w, r)
// }

// func to get all the Books...
func GetBooks(w http.ResponseWriter, r *http.Request) {
	services.GetBooks(w, r)
}
