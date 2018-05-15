package test

import (
	"fmt"

	"github.com/rtt/Go-Solr"
)

func maiiin() {
	// init a connection
	s, err := solr.Init("localhost", 8983, "user")

	if err != nil {
		fmt.Println(err)
		return
	}

	// build an update document, in this case adding two documents
	f := map[string]interface{}{
		"add": []interface{}{
			map[string]interface{}{"id": 1, "uName": "usr1", "avInit": true, "avCr": false, "avAp": true, "avAu": true, "avQA": false, "avRw": true},
			map[string]interface{}{"id": 2, "uName": "usr2", "avInit": true, "avCr": false, "avAp": false, "avAu": false, "avQA": true, "avRw": true},
			map[string]interface{}{"id": 3, "uName": "usr3", "avInit": false, "avCr": true, "avAp": true, "avAu": true, "avQA": false, "avRw": false},
			map[string]interface{}{"id": 4, "uName": "usr4", "avInit": true, "avCr": false, "avAp": false, "avAu": false, "avQA": true, "avRw": true},
			map[string]interface{}{"id": 5, "uName": "usr5", "avInit": false, "avCr": true, "avAp": true, "avAu": true, "avQA": true, "avRw": true},
			map[string]interface{}{"id": 6, "uName": "usr6", "avInit": true, "avCr": true, "avAp": false, "avAu": false, "avQA": false, "avRw": false},
			map[string]interface{}{"id": 7, "uName": "usr7", "avInit": false, "avCr": false, "avAp": true, "avAu": true, "avQA": true, "avRw": true},
			map[string]interface{}{"id": 8, "uName": "usr8", "avInit": true, "avCr": true, "avAp": true, "avAu": true, "avQA": false, "avRw": false},
			map[string]interface{}{"id": 9, "uName": "usr9", "avInit": false, "avCr": false, "avAp": false, "avAu": true, "avQA": false, "avRw": false},
			map[string]interface{}{"id": 10, "uName": "usr10", "avInit": true, "avCr": true, "avAp": false, "avAu": true, "avQA": true, "avRw": true},
			map[string]interface{}{"id": 11, "uName": "usr11", "avInit": false, "avCr": true, "avAp": true, "avAu": false, "avQA": false, "avRw": false},
			map[string]interface{}{"id": 12, "uName": "usr12", "avInit": true, "avCr": false, "avAp": false, "avAu": true, "avQA": false, "avRw": true},
			map[string]interface{}{"id": 13, "uName": "usr13", "avInit": false, "avCr": true, "avAp": true, "avAu": false, "avQA": true, "avRw": false},
			map[string]interface{}{"id": 14, "uName": "usr14", "avInit": true, "avCr": true, "avAp": false, "avAu": true, "avQA": false, "avRw": true},
			map[string]interface{}{"id": 15, "uName": "usr15", "avInit": false, "avCr": false, "avAp": false, "avAu": false, "avQA": true, "avRw": true},
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
