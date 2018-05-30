package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/ninadakolekar/go-dms/src/models"

	"github.com/ninadakolekar/go-dms/src/auth"

	"github.com/gorilla/mux"
	"github.com/ninadakolekar/go-dms/src/docs"
)

// DocViewDetails ... View document header details
func DocViewDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		// User Auth Verification

		username, err := auth.GetCurrentUser(r)

		if err != nil { // Auth unsucessful
			fmt.Println("ERROR docView Line 24: ", err) // Debug
			http.Redirect(w, r, "/", 302)
			return
		}

		// Document View

		vars := mux.Vars(r)
		id := vars["id"]

		if docs.ValidateDocNo(id) {

			Document, err := docs.FetchDocByID(id)

			if err != nil {

				fmt.Println("Failed to fetch document: ", err)
				return
			}

			if Document.FlowStatus == 0 {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			}

			isQA := (username == Document.QA) && (Document.FlowStatus == 1)

			tmpl := template.Must(template.ParseFiles("templates/viewDocDetails.html"))
			tmpl.Execute(w, ViewDetailsMsg{Document, isQA})

		} else {

			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

		}

	}
}

type ViewDetailsMsg struct {
	Doc models.InactiveDoc
	QA  bool
}
