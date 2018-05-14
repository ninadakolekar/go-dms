package main

import (
	"net/http"

	"github.com/gorilla/mux"

	constant "github.com/ninadakolekar/aizant-dms/src/constants"
	controller "github.com/ninadakolekar/aizant-dms/src/controllers"
)

func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/", controller.Index)
	http.ListenAndServe(constant.ApplicationPort, router)
}
