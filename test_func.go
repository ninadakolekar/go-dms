package main

import (
	"net/http"

	constant "github.com/ninadakolekar/aizant-dms/src/constants"
	router "github.com/ninadakolekar/aizant-dms/src/routes"
)

func main() {
	r := router.GetRouter()
	http.ListenAndServe(constant.ApplicationPort, r)
}
