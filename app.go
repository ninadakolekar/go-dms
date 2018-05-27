package main

import (
	"net/http"

	constants "github.com/ninadakolekar/aizant-dms/src/constants"
	router "github.com/ninadakolekar/aizant-dms/src/routes"
)

func main() {

	r := router.GetRouter()
	http.ListenAndServe(constants.ApplicationPort, r)
}
