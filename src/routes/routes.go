package routes

import (
	"github.com/gorilla/mux"
	controller "github.com/ninadakolekar/aizant-dms/src/controllers"
)

// GetRouter ... Returns a new router
func GetRouter() *mux.Router {

	r := mux.NewRouter() // New mux router instance

	// Define routes here
	r.HandleFunc("/hello", controller.Index).Methods("GET")
	return r
}
