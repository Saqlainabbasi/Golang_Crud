package models

import "gorm.io/gorm"

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
