package services

import (
	"github.com/Saqlainabbasi/api/datastructs"
	"github.com/Saqlainabbasi/api/dto"
	"github.com/Saqlainabbasi/api/models"
)

//imports

// interface
type UserService interface {
	CreateUser(user *dto.User) (datastructs.User, error)
	GetUsers() ([]datastructs.User, error)
}

// struct
type userService struct{}

// constructor func
func NewUserService() UserService {
	return &userService{}
}

var (
	UserQuery models.UserQuery = models.NewUserQuery()
)

// interface inplemetation functions
func (*userService) CreateUser(user *dto.User) (datastructs.User, error) {
	// check for any validations....

	//pass the data to the model or repo layer....
	resp, newUser := UserQuery.CreateUser(user)
	if resp.Error != nil {
		return datastructs.User{}, resp.Error
	}
	return newUser, nil
}

// get all user service....
func (*userService) GetUsers() ([]datastructs.User, error) {
	// validations
	//pass to repo layer...

	resp, newUsers := UserQuery.GetUsers()
	if resp.Error != nil {
		return []datastructs.User{}, resp.Error
	}
	return newUsers, nil

}

// import (
// 	"net/http"

// 	"github.com/Saqlainabbasi/api/models"
// 	"github.com/Saqlainabbasi/api/utils"
// )

// func CreateUser(w http.ResponseWriter, r *http.Request) {
// 	//mapping the json data to the DB model format....
// 	// variable of type Struct
// 	user := &models.User{}
// 	// use the utils function formate the data
// 	utils.PraseBody(r, user)
// 	//pass data to database layer to store in db
// 	resp := user.CreateUser()
// 	if resp.Error != nil {
// 		utils.WriteJson(w, http.StatusInternalServerError, resp.Error)
// 	}
// 	// send the responce back
// 	utils.WriteJson(w, http.StatusOK, user)
// }

// // service function to get all the users
// func GetUsers(w http.ResponseWriter, r *http.Request) {
// 	//pass data to database layer to store in db
// 	resp, data := models.GetUsers()
// 	if resp.Error != nil {
// 		utils.WriteJson(w, http.StatusInternalServerError, resp.Error)
// 	}
// 	// send the responce back
// 	utils.WriteJson(w, http.StatusOK, data)
// }
