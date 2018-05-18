package main

import (
	"net/http"

	constant "github.com/ninadakolekar/aizant-dms/src/constants"
	doc "github.com/ninadakolekar/aizant-dms/src/docs"
	router "github.com/ninadakolekar/aizant-dms/src/routes"
)

func main() {
	doc.DeleteInactiveDoc("A-125")
	r := router.GetRouter()
	http.ListenAndServe(constant.ApplicationPort, r)
}
