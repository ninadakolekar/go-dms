package controllers

import (
	"fmt" // Debug
	"html"
	"net/http"

	constant "github.com/ninadakolekar/go-dms/src/constants"
	solr "github.com/rtt/Go-Solr"
)

// DocAvailable ... Checks if a document name is available
func DocAvailable(w http.ResponseWriter, r *http.Request) {

	DocNo := html.EscapeString(r.FormValue("docNumber"))

	// Initialize a solr connection
	s, err := solr.Init(constant.SolrHost, constant.SolrPort, constant.DocsCore)

	if err != nil { // If connection fails
		fmt.Println("ERROR DocAvailable Line 25 ( Solr Connection Failed): ", err) // Debug
		return
	}

	queryString := "id:" + DocNo

	q := solr.Query{
		Params: solr.URLParamMap{
			"q": []string{queryString},
		},
		Rows: 1,
	}

	res, err := s.Select(&q)

	if err != nil {
		fmt.Println("ERROR DocAvailable Line 41 (Failed to query Solr): ", err) // Debug
		return
	}

	results := res.Results

	if results.Len() == 1 {
		// Return false
		fmt.Fprintf(w, "<span> Document Number already exists. </span>")
		return
	}

	// Return True
	fmt.Fprintf(w, "<span> Document number valid. </span>")
}
