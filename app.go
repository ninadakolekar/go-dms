package main

import (
	"net/http"

	constants "github.com/ninadakolekar/go-dms/src/constants"
	router "github.com/ninadakolekar/go-dms/src/routes"
)

func main() {
	r := router.GetRouter()
	http.ListenAndServe(constants.ApplicationPort, r)
}
