package services

import (
	"github.com/Saqlainabbasi/api/dto"
	"github.com/Saqlainabbasi/api/models"
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
	resp, newUser := a.dao.NewUserQuery().CreateUser(user)
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
		return dto.User{}, nil
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
