package routes

import (
	"github.com/Saqlainabbasi/api/middlewares"
	"github.com/gorilla/mux"
)

// initialaze middleware
var Middleware middlewares.AuthMiddleware = middlewares.NewMiddleware()

func CombineRoutes(r *mux.Router) {
	RegisterAuthRoutes(r)
	RegisterUserRoutes(r)
	RegisterBookRoutes(r)
}
