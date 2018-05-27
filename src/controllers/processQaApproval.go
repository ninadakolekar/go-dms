package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	auth "github.com/ninadakolekar/aizant-dms/src/auth"
	constants "github.com/ninadakolekar/aizant-dms/src/constants"
	"github.com/ninadakolekar/aizant-dms/src/docs"
)

// ProcessQaApproval ... Handles QA approval
func ProcessQaApproval(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		// User Auth Verification

		username, err := auth.GetCurrentUser(r)

		if err != nil { // Auth unsucessful
			fmt.Println("ERROR docView Line 24: ", err) // Debug
			http.Redirect(w, r, "/", 302)
		}

		vars := mux.Vars(r)
		id := vars["id"]

		if docs.ValidateDocNo(id) {

			Document, err := docs.FetchDocByID(id)

			if err != nil {

				fmt.Println("Failed to fetch document: ", err)
				return
			}

			if username != Document.QA || Document.FlowStatus != 1 {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			}

			qaResponse := r.FormValue("qa-answer")

			if qaResponse == "approve" {
				Document.FlowStatus = constants.CreateFlow
				docs.AddInactiveDoc(Document)
				fmt.Fprintf(w, "Approval Recorded.")
			} else if qaResponse == "reject" {
				Document.FlowStatus = constants.InitFlow
				docs.AddInactiveDoc(Document)
				fmt.Fprintf(w, "Rejection Recorded.")
			} else {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			}

		}

	} else {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

}
