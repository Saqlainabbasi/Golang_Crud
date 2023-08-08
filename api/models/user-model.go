package models

import (
	"gorm.io/gorm"
)

// declare a struct for user model
// gorm.model is the basic golang staruct,(ID,CreatedAt,....)
type User struct {
	gorm.Model
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    string `json:"age"`
	Email  string `json:"email"`
	Paswrd string `json:"paswrd"`
	Phone  string `json:"phone"`
}

// method to create the user in the database
// this return the user back to the controller...
// this is a reciever fuction...means it receives a Type
// this can only be called with the intence of the Struct
func (u *User) CreateUser() *gorm.DB {
	resp := db.Create(&u)
	return resp
}

// function to get all users
func GetUsers() (*gorm.DB, []User) {
	// []Book means slice of books...
	var user []User
	resp := db.Find(&user)
	return resp, user
}
