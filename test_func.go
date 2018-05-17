package main

import (
	"net/http"

	router "github.com/ninadakolekar/aizant-dms/src/routes"
	constant 
)

func main() {
	r := router.GetRouter()
	http.ListenAndServe(constant.ApplicationPort, r)

}
