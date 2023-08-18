package datastructs

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    string `json:"age"`
	Email  string `json:"email"`
	Paswrd string `json:"paswrd"`
	Phone  string `json:"phone"`
}
