package models

import (
	"github.com/Saqlainabbasi/api/datastructs"
	"gorm.io/gorm"
)

// make interface
type BookQuery interface {
	GetBooks() (*gorm.DB, []datastructs.Book)
	CreateBook(book *datastructs.Book) (*gorm.DB, datastructs.Book)
}

// declare a struct for user model
type bookQuery struct{}

// constructor func
func NewBookQuery() BookQuery {
	return &bookQuery{}
}

// ******************* interface implementation **********************//

// func to get All Books
// this will retrun DB response and data
func (*bookQuery) GetBooks() (*gorm.DB, []datastructs.Book) {
	books := []datastructs.Book{}
	// resp := db.Joins("left join users on books.user_id = users.id").Find(&books)
	resp := db.Find(&books).Joins("users")

	///use reflect package....
	return resp, books
}

// dunc to create book
// this will return dto of book
func (b *bookQuery) CreateBook(book *datastructs.Book) (*gorm.DB, datastructs.Book) {
	// newBook := &datastructs.Book{
	// 	Name:        book.Name,
	// 	Author:      book.Author,
	// 	Publication: book.Publication,
	// 	UserID:      book.UserID,
	// }
	resp := db.Create(book)
	return resp, *book
}

// func GetBooks() (*gorm.DB, []Books) {
// 	//slice of type Books
// 	var books []Books
// 	resp := db.Find(&books)
// 	return resp, books
// }
