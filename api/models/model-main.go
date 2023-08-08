package models

import (
	"github.com/Saqlainabbasi/api/config"
	"gorm.io/gorm"
)

// declare the db variable of type gorm.DB
var db *gorm.DB

func init() {
	_ = config.ConnectToDB()
	db = config.GetDB()
	// fmt.Println(db)
	db.AutoMigrate(&User{})
}