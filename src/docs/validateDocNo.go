package docs

import (
	"strings"

	constant "github.com/ninadakolekar/aizant-dms/src/constants"
	solr "github.com/rtt/Go-Solr"
)

// ValidateDocNo ... Checks if a document number exists
func ValidateDocNo(DocNo string) bool {

	if len(DocNo) <= 2 || strings.Contains(DocNo, " ") { // If length < 3 or if contains whitespace
		return false
	}

	// Initialize a solr connection
	s, err := solr.Init(constant.SolrHost, constant.SolrPort, constant.DocsCore)

	if err != nil { // If connection fails
		return false
	}

	queryString := "id:" + DocNo

	q := solr.Query{
		Params: solr.URLParamMap{
			"q": []string{queryString},
		},
		Rows: 1,
	}

	res, err := s.Select(&q)

	if err != nil { //Failed to query Solr
		return false
	}

	results := res.Results

	if results.Len() == 1 {
		return true
	}

	return false

}
