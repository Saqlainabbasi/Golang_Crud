package routes

import (
	"github.com/gorilla/mux"
)

func CombineRoutes(r *mux.Router) {
	RegisterAuthRoutes(r)
	RegisterUserRoutes(r)
	RegisterBookRoutes(r)
}
