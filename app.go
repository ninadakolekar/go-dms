package main

import (
	"net/http"

	constants "github.com/ninadakolekar/aizant-dms/src/constants"
	router "github.com/ninadakolekar/aizant-dms/src/routes"
	"github.com/ninadakolekar/aizant-dms/test"
)

func main() {
	test.DeleteDocs([]string{"cs16btech11029"})
	r := router.GetRouter()
	http.ListenAndServe(constants.ApplicationPort, r)
}
