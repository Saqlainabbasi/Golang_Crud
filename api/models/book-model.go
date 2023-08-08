package models

import "gorm.io/gorm"

// declare a struct for user model
type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// func to get All Books
// this will retrun DB response and data
func GetBooks() (*gorm.DB, []Book) {
	//slice of type Book
	var books []Book
	resp := db.Find(&books)
	return resp, books
}
