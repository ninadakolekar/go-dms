package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ninadakolekar/aizant-dms/src/docs"
)

// DocAddEdit ... Re-initiating a document
func DocAddEdit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		vars := mux.Vars(r)
		id := vars["id"]

		if docs.ValidateDocNo(id) {

			document, err := docs.FetchDocByID(id)

			if err != nil {
				fmt.Println("Failed to fetch document: ", err)
				return
			}

			tmpl := template.Must(template.ParseFiles("templates/addNewDoc.html"))
			tmpl.Execute(w, docAddMsg{Datab: false, Errb: false, Datamsg: "hi", Approvers: SendApprovers(), Reviewers: SendReviewers(), Authorisers: SendAuthoriser(), Creators: SendCreators(), DocumentExist: true, Redirect: false, Document: document})

		} else {

			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

		}

	}
}
