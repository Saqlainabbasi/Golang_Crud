package middlewares

import (
	"net/http"

	"github.com/Saqlainabbasi/api/services"
	"github.com/Saqlainabbasi/api/utils"
)

// interface
type AuthMiddleware interface {
	Auth(h http.HandlerFunc) http.HandlerFunc
}

type authMiddleware struct {
	tokenManager services.TokenManager
}

// intialize the token manager....
var (
	TokenManager services.TokenManager = services.NewTokenManager("mysecret")
)

// constructor func
func NewMiddleware() AuthMiddleware {
	return &authMiddleware{tokenManager: TokenManager}
}

// interface implementation....
// this function will verify route for access token
func (m *authMiddleware) Auth(handlerFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//get the access token form the Headers
		tokenString := r.Header.Get("Authorization")
		_, err := m.tokenManager.Parse(tokenString)
		if err != nil {
			// fmt.Println(err)
			utils.WriteJson(w, http.StatusForbidden, err.Error())
			return
		}
		handlerFunction(w, r)
	}
}
