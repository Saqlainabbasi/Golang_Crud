package models

import (
	"github.com/Saqlainabbasi/api/datastructs"
	"gorm.io/gorm"
)

// make interface
type BookQuery interface {
	GetBooks() (*gorm.DB, []datastructs.Books)
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
func (*bookQuery) GetBooks() (*gorm.DB, []datastructs.Books) {
	books := []datastructs.Books{}
	resp := db.Find(&books)
	return resp, books
}

// func GetBooks() (*gorm.DB, []Books) {
// 	//slice of type Books
// 	var books []Books
// 	resp := db.Find(&books)
// 	return resp, books
// }
