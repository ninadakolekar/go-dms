package main

import (
	"net/http"

	constants "github.com/ninadakolekar/go-dms/src/constants"
	router "github.com/ninadakolekar/go-dms/src/routes"
	"github.com/ninadakolekar/go-dms/test"
)

func main() {

	papers := test.DBLPResponse()
	test.InsertUsers(papers)
	test.InsertPDF(papers)
	r := router.GetRouter()
	http.ListenAndServe(constants.ApplicationPort, r)
}
