package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Saqlainabbasi/api/config"
	"github.com/Saqlainabbasi/api/routes"
	"github.com/gorilla/mux"
)

func main() {
	// initialize mux router...
	r := mux.NewRouter()
	//connect to DB
	err := config.ConnectToDB()
	if err != nil {
		log.Fatalln("Connection to the Database Unsuccessful")
	}
	// db := config.GetDB()
	// fmt.Println(db)
	//register the different modules routes...
	routes.CombineRoutes(r)
	//handle for api request...
	http.Handle("/", r)
	fmt.Println("server listining on port 8080")
	http.ListenAndServe("localhost:8080", r)
}
