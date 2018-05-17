package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	constant "github.com/ninadakolekar/aizant-dms/src/constants"
)

// ProcessDocAdd ... Process the form-values and add the document
func ProcessDocAdd(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		r.ParseForm()

		docNumber := r.Form["docNumber"][0]
		docName := r.Form["docName"][0]
		docProcess := r.Form["docProcess"][0]
		docType := r.Form["docType"][0]
		docDept := r.Form["docDept"][0]
		docEffDate := r.Form["docEffDate"][0]
		docExpDate := r.Form["docExpDate"][0]
		docCreator := r.Form["docCreator"][0]
		docAuth := r.Form["docAuth"]
		docReviewers := r.Form["docReviewers"]
		docApprovers := r.Form["docApprovers"]
		fmt.Println("Form Received\n ", docNumber, docName, docProcess, docType, docDept, docEffDate, docExpDate, docCreator, docReviewers, docApprovers, docAuth) // Debug
	}

	// Render a new form
	tmpl := template.Must(template.ParseFiles("templates/addNewDoc.html"))
	tmpl.Execute(w, struct{ s uint }{constant.MinDocNumLen})
}
