package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	constant "github.com/ninadakolekar/aizant-dms/src/constants"
)

// ProcessDocAdd ... Process the form-values and add the document
func ProcessDocAdd(w http.ResponseWriter, r *http.Request) {

	docNumber := r.FormValue("docNumber")

	fmt.Println("Form Received for docNumber ", docNumber) // Debug

	tmpl := template.Must(template.ParseFiles("templates/addNewDoc.html"))
	tmpl.Execute(w, struct{ s uint }{constant.MinDocNumLen})
}
