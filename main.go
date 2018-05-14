package main

import (
	"net/http"

	constant "github.com/ninadakolekar/aizant-dms/src/constants"
	route "github.com/ninadakolekar/aizant-dms/src/routes"
)

func main() {
	router := route.GetRouter()
	http.ListenAndServe(constant.ApplicationPort, router)
}
