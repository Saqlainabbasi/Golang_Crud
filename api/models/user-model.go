package models

import (
	"github.com/Saqlainabbasi/api/datastructs"
	"github.com/Saqlainabbasi/api/dto"
	"gorm.io/gorm"
)

type UserQuery interface {
	CreateUser(user *dto.User) (*gorm.DB, datastructs.User)
	GetUsers() (*gorm.DB, []datastructs.User)
	GerUserById(userID int64) (*gorm.DB, datastructs.User)
	GetUserByEmail(email string) (*gorm.DB, dto.User)
}

type userQuery struct{}

// constructor.......
func NewUserQuery() UserQuery {
	return &userQuery{}
}

// interface inplementation......
func (*userQuery) CreateUser(user *dto.User) (*gorm.DB, datastructs.User) {
	newUser := &datastructs.User{
		Name:   user.Name,
		Age:    user.Age,
		Gender: user.Gender,
		Phone:  user.Phone,
		Paswrd: user.Paswrd,
		Email:  user.Email,
	}
	resp := db.Create(newUser)
	return resp, *newUser
}

func (*userQuery) GetUsers() (*gorm.DB, []datastructs.User) {
	//slice of user
	var users = []datastructs.User{}
	resp := db.Find(&users)
	return resp, users
}

func (*userQuery) GerUserById(userId int64) (*gorm.DB, datastructs.User) {
	user := datastructs.User{}
	resp := db.First(&user, userId)
	return resp, user
}

func (*userQuery) GetUserByEmail(email string) (*gorm.DB, dto.User) {
	var user dto.User
	resp := db.Where("email=?", email).First(&user)
	return resp, user
}

// // declare a struct for user model
// // gorm.model is the basic golang staruct,(ID,CreatedAt,....)
// type User struct {
// 	gorm.Model
// 	Name   string `json:"name"`
// 	Gender string `json:"gender"`
// 	Age    string `json:"age"`
// 	Email  string `json:"email"`
// 	Paswrd string `json:"paswrd"`
// 	Phone  string `json:"phone"`
// }

// // method to create the user in the database
// // this return the user back to the controller...
// // this is a reciever fuction...means it receives a Type
// // this can only be called with the intence of the Struct
// func (u *User) CreateUser() *gorm.DB {
// 	resp := db.Create(&u)
// 	return resp
// }

// // function to get all users
// func GetUsers() (*gorm.DB, []User) {
// 	// []Book means slice of books...
// 	var user []User
// 	resp := db.Find(&user)
// 	return resp, user
// }
