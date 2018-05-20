package controllers

import (
	"html/template"
	"net/http"
)

//DocAdd ...
func DocAdd(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/addNewDoc.html"))
	tmpl.Execute(w, templateData{Datab: false, Errb: false, Datamsg: "hi", Approvers: SendApprovers(), Reviewers: SendReviewers(), Authorisers: SendAuthoriser(), Creators: SendCreators()})
}

type templateData struct {
	Datab       bool
	Errb        bool
	Datamsg     string
	Approvers   []string
	Reviewers   []string
	Authorisers []string
	Creators    []string
}
