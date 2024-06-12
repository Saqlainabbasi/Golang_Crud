package controllers

import (
	"net/http"

	"github.com/Saqlainabbasi/api/dto"
	"github.com/Saqlainabbasi/api/services"
	"github.com/Saqlainabbasi/api/utils"
)

// func CreateBook(w http.ResponseWriter, r *http.Request) {
// 	// passing data to service layer for processing
// 	services.CreateUser(w, r)
// }

// initializing the book service
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
	// books := []datastructs.Books{}
	books, err := BookService.GetBooks()
	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, err)
	}
	utils.WriteJson(w, http.StatusOK, books)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value("UID").(int64)

	book := &dto.Book{}

	utils.PraseBody(r, book)
	if book.UserID != 0 && book.UserID != userID {
		utils.WriteJson(w, http.StatusInternalServerError, "Access Denied...! Invalid Data")
		return
	}
	newBook, err := BookService.CreateBook(book)

	if err != nil {
		utils.WriteJson(w, http.StatusInternalServerError, err)
	}
	//sending response back.....
	utils.WriteJson(w, http.StatusOK, newBook)
}
