package test

import (
	"fmt"

	constants "github.com/ninadakolekar/aizant-dms/src/constants"
	"github.com/rtt/Go-Solr"
)

//AddUsers ... addsuser
func AddUsers() {
	// init a connection
	s, err := solr.Init(constants.SolrHost, constants.SolrPort, constants.UserCore)

	if err != nil {
		fmt.Println(err)
		return
	}

	// build an update document, in this case adding two documents
	f := map[string]interface{}{
		"add": []interface{}{
			map[string]interface{}{"id": "dinesh1997", "uName": "Dinesh Tripathi", "avInit": true, "avCr": false, "avAp": true, "avAu": true, "avQA": false, "avRw": true},
			map[string]interface{}{"id": "adityaM", "uName": "Aditya Menon", "avInit": true, "avCr": false, "avAp": false, "avAu": false, "avQA": false, "avRw": true},
			map[string]interface{}{"id": "ssingh", "uName": "Suresh Singh", "avInit": false, "avCr": true, "avAp": true, "avAu": true, "avQA": false, "avRw": false},
			map[string]interface{}{"id": "agarh", "uName": "Anil Garh", "avInit": true, "avCr": false, "avAp": false, "avAu": false, "avQA": false, "avRw": true},
			map[string]interface{}{"id": "baldev", "uName": "Baldev Sodhani", "avInit": false, "avCr": true, "avAp": true, "avAu": true, "avQA": false, "avRw": true},
			map[string]interface{}{"id": "klein", "uName": "Wilburn Klein", "avInit": true, "avCr": true, "avAp": false, "avAu": false, "avQA": false, "avRw": false},
			map[string]interface{}{"id": "akofod", "uName": "Andreas Kofod", "avInit": false, "avCr": false, "avAp": true, "avAu": true, "avQA": false, "avRw": true},
			map[string]interface{}{"id": "ooriee", "uName": "Orie McLaughlin", "avInit": true, "avCr": true, "avAp": true, "avAu": true, "avQA": false, "avRw": false},
			map[string]interface{}{"id": "gregH", "uName": "Greg Hendricks", "avInit": false, "avCr": false, "avAp": false, "avAu": true, "avQA": false, "avRw": false},
			map[string]interface{}{"id": "Gugu", "uName": "Gugulethu", "avInit": true, "avCr": true, "avAp": false, "avAu": true, "avQA": false, "avRw": true},
			map[string]interface{}{"id": "irene", "uName": "Ireneport", "avInit": false, "avCr": true, "avAp": true, "avAu": false, "avQA": false, "avRw": false},
			map[string]interface{}{"id": "peter", "uName": "Petersen", "avInit": true, "avCr": false, "avAp": false, "avAu": true, "avQA": false, "avRw": true},
			map[string]interface{}{"id": "atque", "uName": "Atque", "avInit": false, "avCr": true, "avAp": true, "avAu": false, "avQA": false, "avRw": false},
			map[string]interface{}{"id": "nicol", "uName": "Nicolette De Bruin", "avInit": true, "avCr": true, "avAp": false, "avAu": true, "avQA": false, "avRw": true},
			map[string]interface{}{"id": "firefox", "uName": "Mozilla", "avInit": false, "avCr": false, "avAp": false, "avAu": false, "avQA": true, "avRw": true},
		},
	}

	// send off the update (2nd parameter indicates we also want to commit the operation)
	resp, err := s.Update(f, true)

	if err != nil {
		fmt.Println("error =>", err)
	} else {
		fmt.Println("resp =>", resp)
	}
}
