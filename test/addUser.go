package test

import (
	"fmt"

	"github.com/rtt/Go-Solr"
)

//AddUser ... addsuser
func AddUser() {
	// init a connection
	s, err := solr.Init("localhost", 8983, "user")

	if err != nil {
		fmt.Println(err)
		return
	}

	// build an update document, in this case adding two documents
	f := map[string]interface{}{
		"add": []interface{}{
			map[string]interface{}{"id": "user1", "uName": "Dinesh Lal Tripathi", "avInit": true, "avCr": false, "avAp": true, "avAu": true, "avQA": false, "avRw": true},
			map[string]interface{}{"id": "user2", "uName": "Aditya Menon", "avInit": true, "avCr": false, "avAp": false, "avAu": false, "avQA": true, "avRw": true},
			map[string]interface{}{"id": "user3", "uName": "Obaid Lal Suresh", "avInit": false, "avCr": true, "avAp": true, "avAu": true, "avQA": false, "avRw": false},
			map[string]interface{}{"id": "user4", "uName": "AnilGarh", "avInit": true, "avCr": false, "avAp": false, "avAu": false, "avQA": true, "avRw": true},
			map[string]interface{}{"id": "user5", "uName": "Baldev Sodhani", "avInit": false, "avCr": true, "avAp": true, "avAu": true, "avQA": true, "avRw": true},
			map[string]interface{}{"id": "user6", "uName": "Wilburn Klein", "avInit": true, "avCr": true, "avAp": false, "avAu": false, "avQA": false, "avRw": false},
			map[string]interface{}{"id": "user7", "uName": "Andreas Kofod", "avInit": false, "avCr": false, "avAp": true, "avAu": true, "avQA": true, "avRw": true},
			map[string]interface{}{"id": "user8", "uName": "Orie McLaughlin", "avInit": true, "avCr": true, "avAp": true, "avAu": true, "avQA": false, "avRw": false},
			map[string]interface{}{"id": "user9", "uName": "Greg Hendricks", "avInit": false, "avCr": false, "avAp": false, "avAu": true, "avQA": false, "avRw": false},
			map[string]interface{}{"id": "user10", "uName": "Gugulethu", "avInit": true, "avCr": true, "avAp": false, "avAu": true, "avQA": true, "avRw": true},
			map[string]interface{}{"id": "user11", "uName": "Ireneport", "avInit": false, "avCr": true, "avAp": true, "avAu": false, "avQA": false, "avRw": false},
			map[string]interface{}{"id": "user12", "uName": "Petersen", "avInit": true, "avCr": false, "avAp": false, "avAu": true, "avQA": false, "avRw": true},
			map[string]interface{}{"id": "user13", "uName": "atque", "avInit": false, "avCr": true, "avAp": true, "avAu": false, "avQA": true, "avRw": false},
			map[string]interface{}{"id": "user14", "uName": "Nicolette De Bruin", "avInit": true, "avCr": true, "avAp": false, "avAu": true, "avQA": false, "avRw": true},
			map[string]interface{}{"id": "user15", "uName": "Mozilla", "avInit": false, "avCr": false, "avAp": false, "avAu": false, "avQA": true, "avRw": true},
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
