package main

import (
	"net/http"

	"github.com/gorilla/mux"
	controller "github.com/ninadakolekar/aizant-dms/src/controllers"
	test "github.com/ninadakolekar/aizant-dms/test"
)

func main() {
	r := mux.NewRouter() // New mux router instance

	// Serve Static Files
	r.PathPrefix("/static/css").Handler(http.StripPrefix("/static/css/", http.FileServer(http.Dir("static/css/"))))
	r.PathPrefix("/static/js").Handler(http.StripPrefix("/static/js/", http.FileServer(http.Dir("static/js/"))))
	r.PathPrefix("/static/images").Handler(http.StripPrefix("/static/images/", http.FileServer(http.Dir("static/images/"))))

	// Define routes here
	r.HandleFunc("/hello", controller.Index).Methods("GET")
	r.HandleFunc("/users", controller.Users).Methods("GET")
	r.HandleFunc("/doc/add", controller.DocAdd).Methods("GET")
	r.HandleFunc("/doc/avail", controller.DocAvailable).Methods("POST")

	http.ListenAndServe(":8080", r)
	test.TestAddDelDoc()
}