package services

import (
	"github.com/Saqlainabbasi/api/datastructs"
	"github.com/Saqlainabbasi/api/dto"
	"github.com/Saqlainabbasi/api/models"
	"github.com/Saqlainabbasi/api/utils"
	"github.com/jinzhu/copier"
)

// interface
type AuthService interface {
	SignUp(user *dto.User) (dto.User, error)
	SignIn(email, password string) (dto.User, error)
	Logout(userId int64) error
}

// declare struct
type authService struct {
	dao          models.DAO
	tokenManager TokenManager
}

// constrctor function
func NewAuthService(dao models.DAO, tokenManager TokenManager) AuthService {
	return &authService{dao: dao, tokenManager: tokenManager}
}

// implementation interface
func (a *authService) SignUp(user *dto.User) (dto.User, error) {
	//convert the dto to the datastruct...
	_user := &datastructs.User{}
	utils.DataMapper(&_user, &user, copier.Option{IgnoreEmpty: true})
	resp, newUser := a.dao.NewUserQuery().CreateUser(_user)
	if resp.Error != nil {
		return dto.User{}, resp.Error
	}
	data := dto.User{
		ID:     newUser.ID,
		Name:   newUser.Name,
		Gender: newUser.Gender,
		Age:    newUser.Age,
		Email:  newUser.Email,
		Phone:  newUser.Phone,
	}
	return data, nil
}

func (a *authService) SignIn(email, password string) (dto.User, error) {
	//get user by email...
	// var user dto.User
	resp, user := a.dao.NewUserQuery().GetUserByEmail(email)
	if resp.Error != nil {
		return dto.User{}, resp.Error
	}
	//check user passord...

	//generate the jwt

	jwt, err := a.tokenManager.NewJWT(int64(user.ID))
	if err != nil {
		return dto.User{}, err
	}
	user.Token = jwt
	return user, nil
}

func (a *authService) Logout(userId int64) error {
	return nil
}
