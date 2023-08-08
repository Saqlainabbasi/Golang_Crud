package config

import (
	"fmt"

	// "github.com/Saqlainabbasi/api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// declare a variable of type gorm
var DB *gorm.DB

// function to connect to the Database
func ConnectToDB() error {

	//connection string....
	dns := "root:testDB@tcp(127.0.0.1:3306)/myDB?charset=utf8mb4&parseTime=True&loc=Local"
	// send the requset to database using gorm...
	con, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		// log.Fatal("Error")
		fmt.Println("Error Connecting DB")
		return err
	}
	fmt.Println("Connection Successful")
	// con.AutoMigrate(&models.User{})
	DB = con
	return nil
}

func GetDB() *gorm.DB {
	return DB
}
