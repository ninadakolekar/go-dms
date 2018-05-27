package main

import (
	"net/http"

	"github.com/ninadakolekar/aizant-dms/src/docs"

	constants "github.com/ninadakolekar/aizant-dms/src/constants"
	router "github.com/ninadakolekar/aizant-dms/src/routes"
)

func main() {
	doc, _ := docs.FetchDocByID("CS3523")
	doc.DocProcess = "OneByOne"
	docs.AddInactiveDoc(doc)
	r := router.GetRouter()
	http.ListenAndServe(constants.ApplicationPort, r)
}
