package routes

import (
	"github.com/Saqlainabbasi/api/controllers"
	"github.com/gorilla/mux"
)

// this function will handle all user routes...
// takes a pointer of *mux.Router
func RegisterAuthRoutes(router *mux.Router) {
	router.HandleFunc("/signUp/", controllers.SignUp).Methods("POST")
	router.HandleFunc("/signIn/", controllers.SignIn).Methods("POST")
}

// //interface
// type AuthRoutes interface{
// 	RegisterAuthRoutes()
// }
// //struce
// type authRoute struct{}
// //constructor
// func NewAuthRoutes() AuthRoutes{
// 	return &authRoute{}
// }
// //implementation
// func (ar *authRoute) RegisterAuthRoutes(){
// 	ar.HandleFunc("/signUp/", controllers.SignUp).Methods("POST")
// }
