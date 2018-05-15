package test

import (
	"fmt"

	doc "github.com/ninadakolekar/aizant-dms/src/docs"
	model "github.com/ninadakolekar/aizant-dms/src/models"
)

func maiin() {
	x := model.InactiveDoc{"A-23", "Pharma Pract", "SOP", false, "Ramki", "Ramki", []string{"Ramki"}, []string{"Ramki"}, []string{"Ramki"}, "HR", 0, 0, "Body of the Pharma Doc"}
	resp, err := doc.AddInactiveDoc(x)
	if err != nil {
		fmt.Println("ERROR main() Line 13: " + err.Error())
	} else {
		fmt.Println(resp)
	}
}
