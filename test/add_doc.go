package main

import (
	"fmt"

	"github.com/rtt/Go-Solr"
)

func main() {
	s, err := solr.Init("localhost", 8983, "docs")

	if err != nil {
		fmt.Println(err)
		return
	}

	// build an update document, in this case adding two documents
	f := map[string]interface{}{
		"add": []interface{}{
			map[string]interface{}{"id": 22,
				"DocNo":         "402",
				"title":         "abc",
				"docStatus":     false,
				"approver":      []string{"ramki", "kishan"},
				"authorizer":    "ninad",
				"creator":       "shanmuck",
				"initiator":     "kishan",
				"docDepartment": "cse",
				"docTemplateID": 1,
				"body":          "lorem ipsum",
				"docType":       "test",
				"flowStatus":    5,
				"reviewer":      "kishan",
			},
			map[string]interface{}{"id": 12,
				"DocNo":         "42",
				"title":         "ab dc",
				"docStatus":     false,
				"approver":      []string{"ramki1", "kishan1"},
				"authorizer":    "ninad1",
				"creator":       "shanmuck1",
				"initiator":     "kishan1",
				"docDepartment": "cse1",
				"docTemplateID": 1,
				"body":          "lorem ipsum 1",
				"docType":       "test1",
				"flowStatus":    51,
				"reviewer":      "kishan1",
			},
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
