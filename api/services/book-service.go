package services

import (
	"github.com/Saqlainabbasi/api/datastructs"
	"github.com/Saqlainabbasi/api/models"
)

// make interface....
type BookService interface {
	GetBooks() ([]datastructs.Books, error)
}

// declare struct....
type bookService struct{}

// constructor func.....
func NewBookService() BookService {
	return &bookService{}
}

var (
	BookQuery models.BookQuery = models.NewBookQuery()
)

// interface implemetation...

// func to get all books
func (*bookService) GetBooks() ([]datastructs.Books, error) {
	//validations checks

	// call repo/model layer
	resp, newBooks := BookQuery.GetBooks()
	if resp.Error != nil {
		return []datastructs.Books{}, resp.Error
	}
	return newBooks, nil
}

// func GetBooks(w http.ResponseWriter, r *http.Request) {
// 	//pass to Database layer
// 	resp, data := models.GetBooks()
// 	if resp.Error != nil {
// 		utils.WriteJson(w, http.StatusInternalServerError, resp.Error)
// 	}
// 	// send the responce back
// 	utils.WriteJson(w, http.StatusOK, data)
// }
