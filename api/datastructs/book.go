package datastructs

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
	UserID      int64  `json:"user_id"`
}
