package controllers

import (
	"fmt" // Debug
	"net/http"

	constant "github.com/ninadakolekar/aizant-dms/src/constants"
	solr "github.com/rtt/Go-Solr"
)

// DocAvailable ... Checks if a document name is available
func DocAvailable(w http.ResponseWriter, r *http.Request) {
	// Initialize a solr connection
	s, err := solr.Init(constant.SolrHost, constant.SolrPort, constant.DocsCore)

	if err != nil { // If connection fails
		fmt.Println("ERROR AddInactiveDoc Line 14: ", err) // Debug
		return
	}

	fmt.Print(s) // Remove this line

	// Take form data recieved in POST request
	// Query that document number --> If found send false else send true
}
