package routes

import (
	"github.com/Saqlainabbasi/api/controllers"
	"github.com/gorilla/mux"
)

func RegisterBookRoutes(router *mux.Router) {
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
}
