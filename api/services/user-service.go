package services

import (
	"context"

	"github.com/Saqlainabbasi/api/datastructs"
	"github.com/Saqlainabbasi/api/dto"
	"github.com/Saqlainabbasi/api/models"
	"github.com/Saqlainabbasi/api/utils"
	"github.com/jinzhu/copier"
)

//imports

// interface
type UserService interface {
	CreateUser(user *dto.User) (dto.User, error)
	GetUsers() ([]dto.User, error)
}

// struct
type userService struct {
	amqp RabbitMQService
}

// constructor func
func NewUserService(amqp RabbitMQService) UserService {
	return &userService{amqp: amqp}
}

var (
	UserQuery models.UserQuery = models.NewUserQuery()
)

// interface inplemetation functions
func (us *userService) CreateUser(user *dto.User) (dto.User, error) {
	// check for any validations....

	//convert the dto to the datastruct...
	_user := &datastructs.User{}
	utils.DataMapper(&_user, &user, copier.Option{IgnoreEmpty: true})
	//pass the data to the model or repo layer....
	resp, newData := UserQuery.CreateUser(_user)
	if resp.Error != nil {
		return dto.User{}, resp.Error
	}
	newUser := dto.User(*user)
	ctx := context.Background()
	reqBody, err := utils.JsonMarshal(newUser)
	if err != nil {
		return dto.User{}, resp.Error
	}
	us.amqp.Pubblisher(ctx, "create_user_service", reqBody)
	newUser.ID = newData.ID
	return newUser, nil
}

// get all user service....
func (*userService) GetUsers() ([]dto.User, error) {
	// validations
	//pass to repo layer...

	resp, newUsers := UserQuery.GetUsers()
	//map datastruct slice to dto slice
	_users := []dto.User{}
	utils.DataMapper(&_users, &newUsers, copier.Option{})
	if resp.Error != nil {
		return []dto.User{}, resp.Error
	}
	return _users, nil

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
