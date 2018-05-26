package user

import (
	"fmt" // Debug

	constant "github.com/ninadakolekar/aizant-dms/src/constants"
	solr "github.com/rtt/Go-Solr"
)

// DeleteUser ... Deletes auser
func DeleteUser(username string) (*solr.UpdateResponse, error) {

	// Initialize a solr connection
	s, err := solr.Init(constant.SolrHost, constant.SolrPort, constant.UserCore)

	if err != nil { // If connection fails
		fmt.Println("ERROR DeleteUser Line 17: ", err) // Debug
		return nil, err
	}

	// build an update document, in this case adding two documents
	g := map[string]interface{}{
		"delete": []interface{}{
			map[string]interface{}{
				"id": username,
			},
		},
	}

	// send off the update (2nd parameter indicates we also want to commit the operation)
	resp, err := s.Update(g, true)

	if err != nil {
		return resp, err
	}

	return resp, nil
}
