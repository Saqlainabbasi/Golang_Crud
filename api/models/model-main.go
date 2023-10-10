package models

import (
	// "github.com/Saqlainabbasi/api/config"
	// "github.com/Saqlainabbasi/api/datastructs"

	"log"

	"github.com/Saqlainabbasi/api/config"
	"gorm.io/gorm"
)

// DAO interface (Data Access Object)
type DAO interface {
	NewUserQuery() UserQuery
	NewBookQuery() BookQuery
}

// dao struct
type dao struct{}

// declare the db variable of type gorm.DB
var db *gorm.DB

func init() {
	// config_new := config.DBConfig{Type: "mysql"}

	// _ = config_new.ConnectToDB()
	// _ = config.ConnectToDB()
	err := config.ConnectToDB()
	if err != nil {
		log.Fatalf("%s database connection not found", err.Error())
	}
	db = config.GetDB()
	// fmt.Println(db)
	// db.AutoMigrate(&datastructs.User{})
}
func NewDAO() DAO {
	return &dao{}
}

// interface implementaion.....
func (*dao) NewUserQuery() UserQuery {
	return &userQuery{}
}

func (*dao) NewBookQuery() BookQuery {
	return &bookQuery{}
}
