package routes

import (
	"github.com/Saqlainabbasi/api/controllers"
	"github.com/gorilla/mux"
)

// this function will handle all user routes...
// takes a pointer of *mux.Router

func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/user/", Middleware.Auth(controllers.GetUsers)).Methods("GET")
}
