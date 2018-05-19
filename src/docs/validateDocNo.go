package docs

import (
	constant "github.com/ninadakolekar/aizant-dms/src/constants"
	solr "github.com/rtt/Go-Solr"
)

// ValidateDocNo ... Checks if a document number exists
// func ValidateDocNo(str string) bool {
// 	s, err := solr.Init(constant.SolrHost, constant.SolrPort, constant.DocsCore)

// 	if err != nil {
// 		fmt.Println(err)
// 		return false
// 	}
// 	quer := "id:" + str
// 	// fmt.Println("query : ", quer)
// 	q := solr.Query{ //checking in backend whether any other documnet with same id is present

// 		Params: solr.URLParamMap{
// 			"q": []string{quer},
// 		},
// 		Rows: 1,
// 	}
// 	res, err := s.Select(&q)
// 	if err != nil {
// 		fmt.Println(err)
// 		return false
// 	}

// 	results := res.Results
// 	if results.Len() != 0 {
// 		return false
// 	}
// 	return true
// }

func ValidateDocNo(DocNo string) bool {

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
