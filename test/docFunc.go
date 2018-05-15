package test

import (
	"fmt"

	doc "github.com/ninadakolekar/aizant-dms/src/docs"
	model "github.com/ninadakolekar/aizant-dms/src/models"
)

func TestAddDelDoc() {
	x := model.InactiveDoc{"A-23", "Pharma Pract", "SOP", false, "Ramki", "Ramki", []string{"Ramki"}, []string{"Ramki"}, []string{"Ramki"}, "HR", 0, 0, "Body of the Pharma Doc"}
	resp, err := doc.AddInactiveDoc(x)
	if err != nil {
		fmt.Println("ERROR main() Line 13: " + err.Error()) // Debug
	} else {
		fmt.Println(resp)
	}

	var i int
	_, err = fmt.Scanf("%d", &i)

	resp, err = doc.DeleteInactiveDoc("A-23")
	if err != nil {
		fmt.Println("ERROR main() Line 21: " + err.Error()) // Debug
	} else {
		fmt.Println(resp)
	}
}
