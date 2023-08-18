package controllers

import (
	"net/http"

	"github.com/Saqlainabbasi/api/datastructs"
	"github.com/Saqlainabbasi/api/services"
	"github.com/Saqlainabbasi/api/utils"
)

// func CreateBook(w http.ResponseWriter, r *http.Request) {
// 	// passing data to service layer for processing
// 	services.CreateUser(w, r)
// }

var (
	BookService services.BookService = services.NewBookService()
)

// // func to get all the Books...
// func GetBooks(w http.ResponseWriter, r *http.Request) {
// 	services.GetBooks(w, r)
// }

// new func with interfaces....BookService
func GetBooks(w http.ResponseWriter, r *http.Request) {
	//slice  books of type dto.Books
	books := []datastructs.Books{}
	books, err := BookService.GetBooks()
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, err)
	}
	utils.WriteJson(w, http.StatusOK, books)

}
