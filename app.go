package main

import (
	"net/http"

	constants "github.com/ninadakolekar/aizant-dms/src/constants"
	router "github.com/ninadakolekar/aizant-dms/src/routes"
	"github.com/ninadakolekar/aizant-dms/test"
)

func main() {
	test.DeleteDocs([]string{"CS1330-ITP", "CS3523-OS2", "CS2420-ICT"})
	r := router.GetRouter()
	http.ListenAndServe(constants.ApplicationPort, r)
}
