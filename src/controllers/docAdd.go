package controllers

import (
	"html/template"
	"net/http"

	"github.com/ninadakolekar/aizant-dms/src/models"
)

//DocAdd ...
func DocAdd(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/addNewDoc.html"))
	tmpl.Execute(w, docAddMsg{Datab: false, Errb: false, Datamsg: "hi", Approvers: SendApprovers(), Reviewers: SendReviewers(), Authorisers: SendAuthoriser(), Creators: SendCreators(), DocumentExist: false, Redirect: false, Document: models.InactiveDoc{}})
}

type docAddMsg struct {
	Datab         bool
	Errb          bool
	Datamsg       string
	Approvers     []string
	Reviewers     []string
	Authorisers   []string
	Creators      []string
	DocumentExist bool
	Redirect      bool
	Document      models.InactiveDoc
}
