package user

import (
	"fmt"

	constants "github.com/ninadakolekar/aizant-dms/src/constants"
	"github.com/ninadakolekar/aizant-dms/src/models"
	"github.com/rtt/Go-Solr"
)

//AddUser ... Adds a given user to solr database

func AddUser(user models.User) (*solr.UpdateResponse, error) {

	// Initialize a Solr connection
	s, err := solr.Init(constants.SolrHost, constants.SolrPort, constants.UserCore)

	if err != nil {
		fmt.Println("ERROR addUser Line 19: ", err) // Debug
		return nil, err
	}

	// build an update document, in this case adding two documents
	f := map[string]interface{}{
		"add": []interface{}{
			map[string]interface{}{"id": user.Username, "uName": user.Name, "avInit": user.AvailableInit, "avCr": user.AvailableCr, "avAp": user.AvailableApp, "avAu": user.AvailableAuth, "avQA": user.AvailableQA, "avRw": user.AvailableRw},
		},
	}

	// Send off the update (2nd parameter indicates we also want to commit the operation)
	resp, err := s.Update(f, true)

	if err != nil {
		fmt.Println("ERROR addUser Line 34: ", err) // Debug
		return resp, err
	}

	return resp, err
}
