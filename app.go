package main

import (
	"net/http"

	"github.com/ninadakolekar/aizant-dms/src/docs"

	constants "github.com/ninadakolekar/aizant-dms/src/constants"
	router "github.com/ninadakolekar/aizant-dms/src/routes"
)

func main() {
	document, _ := docs.FetchDocByID("CS1330")
	document.FlowStatus = 5
	document.DocStatus = false
	docs.AddInactiveDoc(document)
	r := router.GetRouter()
	http.ListenAndServe(constants.ApplicationPort, r)
}
