package main

import (
	"net/http"

	constants "github.com/ninadakolekar/go-dms/src/constants"
	router "github.com/ninadakolekar/go-dms/src/routes"
	"github.com/ninadakolekar/go-dms/test"
)

func main() {
	//test.ConvertPDF2StringSlice("test")
	//test.DBLPResponse()
	test.ConvertURL2Strings("http://www.vldb.org/conf/2007/papers/research/p411-abadi.pdf", "1")
	//test.ConvertURLtotext("http://www.vldb.org/conf/2007/papers/research/p411-abadi.pdf")
	r := router.GetRouter()
	http.ListenAndServe(constants.ApplicationPort, r)
}
