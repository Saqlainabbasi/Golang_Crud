package services

import (
	"github.com/Saqlainabbasi/api/datastructs"
	"github.com/Saqlainabbasi/api/dto"
	"github.com/Saqlainabbasi/api/models"
	"github.com/Saqlainabbasi/api/utils"
	"github.com/jinzhu/copier"
)

// make interface....
type BookService interface {
	GetBooks() ([]dto.Book, error)
	CreateBook(book *dto.Book) (dto.Book, error)
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
func (b *bookService) GetBooks() ([]dto.Book, error) {
	//validations checks

	// call repo/model layer
	resp, newBooks := BookQuery.GetBooks()
	//Book slice of type Dto.Books
	Books := []dto.Book{}
	// this will (change the datastruct to dto)
	// copier.CopyWithOption(&Books, &newBooks, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	utils.DataMapper(&Books, &newBooks, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if resp.Error != nil {
		return []dto.Book{}, resp.Error
	}
	return Books, nil
}

// dunc to create Book
func (*bookService) CreateBook(book *dto.Book) (dto.Book, error) {
	//validations checks....
	_book := &datastructs.Book{}
	utils.DataMapper(&_book, &book, copier.Option{})
	resp, newBook := BookQuery.CreateBook(_book)
	if resp.Error != nil {
		return dto.Book{}, resp.Error
	}
	data := dto.Book(*book)
	data.ID = newBook.ID

	return data, nil
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
