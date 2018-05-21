package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	controller "github.com/ninadakolekar/aizant-dms/src/controllers"
)

// GetRouter ... Returns a new router
func GetRouter() *mux.Router {

	r := mux.NewRouter() // New mux router instance

	// Serve Static Files
	r.PathPrefix("/static/css").Handler(http.StripPrefix("/static/css/", http.FileServer(http.Dir("static/css/"))))
	r.PathPrefix("/static/js").Handler(http.StripPrefix("/static/js/", http.FileServer(http.Dir("static/js/"))))
	r.PathPrefix("/static/images").Handler(http.StripPrefix("/static/images/", http.FileServer(http.Dir("static/images/"))))

	// Define routes here
	r.HandleFunc("/hello", controller.Index).Methods("GET")

	r.HandleFunc("/doc/avail", controller.DocAvailable).Methods("POST")

	r.HandleFunc("/doc/add", controller.DocAdd).Methods("GET")
	r.HandleFunc("/doc/add", controller.ProcessDocAdd).Methods("POST")

	r.HandleFunc("/doc/add/{id}", controller.DocAddEdit).Methods("GET")
	r.HandleFunc("/doc/add/{id}", controller.ProcessDocAddEdit).Methods("POST")

	r.HandleFunc("/doc/search", controller.DocSearch).Methods("GET")
	r.HandleFunc("/doc/search", controller.ProcessDocSearch).Methods("POST")

	r.HandleFunc("/doc/create/{id}", controller.DocCreate).Methods("GET")
	r.HandleFunc("/doc/create", controller.ProcessDocCreate).Methods("POST")

	r.HandleFunc("/doc/view/{id}", controller.DocView).Methods("GET")

	return r
}
