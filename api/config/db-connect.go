package config

import (
	"fmt"

	// "github.com/Saqlainabbasi/api/models"
	"github.com/Saqlainabbasi/api/datastructs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// declare a variable of type gorm
var DB *gorm.DB

// define and connection interface
type DBConnection interface {
	ConnectToDB() error
}

// define struct  for the Database config variables......// DBConfig represents the configuration for a database.
type DBConfig struct {
	IdentificationName string // IdentificationName is used to obtain the specific database connection.
	DB                 string // Database name.
	User               string // Database user.
	Password           string `json:"_"` // Database password.
	Host               string // Database host.
	Port               string // Database port.
	Type               string // Type of the database ("mysql", "postgres", "mssql", etc.).
	SSLMode            string // SSL mode for the database connection.
	TimeZone           string // Time zone for the database.
}

// function to connect to the Database
// func ConnectToDB() error {

// 	//connection string....
// 	dns := "root:testDB@tcp(127.0.0.1:3306)/myDB?charset=utf8mb4&parseTime=True&loc=Local"
// 	// send the requset to database using gorm...
// 	con, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
// 	if err != nil {
// 		// log.Fatal("Error")
// 		fmt.Println("Error Connecting DB")
// 		return err
// 	}
// 	fmt.Println("Connection Successful")
// 	// con.AutoMigrate(&models.User{})
// 	DB = con
// 	return nil
// }

func ConnectToDB() error {
	// 	//connection  string......
	// check env for Database Type and get the config
	// switch expression {
	// case condition:

	// }
	dns, _ := MySQLConnectionString()

	con, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		// log.Fatal("Error")
		fmt.Println("Error Connecting DB")
		return err
	}
	fmt.Println("Connection Successful")
	// con.AutoMigrate(&models.User{})
	con.AutoMigrate(&datastructs.User{}, &datastructs.Book{})
	DB = con
	return nil
}

func GetDB() *gorm.DB {
	return DB
}
