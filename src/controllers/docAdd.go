package controllers

import (
	"html/template"
	"net/http"
)

//DocAdd ... not competed
func DocAdd(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/addNewDoc.html"))
	tmpl.Execute(w, templateData{Approvers: SendApprovers(), Reviewers: SendReviewers(), Authorisers: SendAuthoriser(), Creators: SendCreators()})
}

type templateData struct {
	Approvers   []string
	Reviewers   []string
	Authorisers []string
	Creators    []string
}
