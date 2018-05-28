package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	auth "github.com/ninadakolekar/aizant-dms/src/auth"
	"github.com/ninadakolekar/aizant-dms/src/constants"
	"github.com/ninadakolekar/aizant-dms/src/docs"
)

// ProcessAuthApproval ... Handles Authoriser's consent
func ProcessAuthApproval(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		// User Auth Verification

		username, err := auth.GetCurrentUser(r)

		if err != nil { // Auth unsucessful
			fmt.Println("ERROR ProcessAuthApproval Line 23: ", err) // Debug
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

			if !docs.CheckCurrentAuthorizer(Document, username) || Document.FlowStatus != constants.AuthFlow {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}

			authResponse := r.FormValue("auth-answer")

			if authResponse == "approve" {

				if Document.DocProcess == constants.Everyone {

					Document.FlowList = append(Document.FlowList, username)

					if len(Document.FlowList) == len(Document.Authorizer) {
						Document.FlowList = nil
						Document.CurrentFlowUser = 0
						Document.FlowStatus = constants.ActiveFlow
						Document.DocStatus = true
					}

					_, err := docs.AddInactiveDoc(Document)

					if err != nil {
						fmt.Println("Error: Document Update Failed ProcessAuthApproval Line 63 ", err) // Debug
						return
					}

					fmt.Fprintf(w, "Approval Recorded.")

				} else if Document.DocProcess == constants.Anyone {

					Document.FlowStatus = constants.ActiveFlow
					Document.DocStatus = true
					Document.FlowList = nil
					Document.CurrentFlowUser = 0

					_, err := docs.AddInactiveDoc(Document)

					if err != nil {
						fmt.Println("Error: Document Update Failed ProcessAuthApproval Line 79 ", err) // Debug
						return
					}

					fmt.Fprintf(w, "Approval Recorded.")

				} else if Document.DocProcess == constants.OneByOne {

					Document.CurrentFlowUser++

					if int(Document.CurrentFlowUser) == len(Document.Approver) {

						Document.FlowStatus = constants.ActiveFlow
						Document.DocStatus = true
						Document.FlowList = nil
						Document.CurrentFlowUser = 0
					}

					_, err := docs.AddInactiveDoc(Document)

					if err != nil {
						fmt.Println("Error: Document Update Failed ProcessAuthApproval Line 100 ", err) // Debug
						return
					}

					fmt.Fprintf(w, "Approval Recorded.")

				}

			} else if authResponse == "reject" {

				if Document.DocProcess == constants.Everyone {

					Document.FlowStatus = constants.CreateFlow
					Document.FlowList = nil
					Document.CurrentFlowUser = 0

					_, err := docs.AddInactiveDoc(Document)

					if err != nil {
						fmt.Println("Error: Document Update Failed ProcessAuthApproval Line 130 ", err) // Debug
						return
					}

					fmt.Fprintf(w, "Rejection Recorded.")

				} else if Document.DocProcess == constants.Anyone {

					Document.FlowList = append(Document.FlowList, username)

					if len(Document.FlowList) == len(Document.Authorizer) {
						Document.FlowStatus = constants.CreateFlow
						Document.FlowList = nil
						Document.CurrentFlowUser = 0
					}

					_, err := docs.AddInactiveDoc(Document)

					if err != nil {
						fmt.Println("Error: Document Update Failed ProcessAuthApproval Line 149 ", err) // Debug
						return
					}

					fmt.Fprintf(w, "Rejection Recorded.")

				} else if Document.DocProcess == constants.OneByOne {

					Document.FlowStatus = constants.CreateFlow
					Document.FlowList = nil
					Document.CurrentFlowUser = 0

					_, err := docs.AddInactiveDoc(Document)

					if err != nil {
						fmt.Println("Error: Document Update Failed ProcessAuthApproval Line 164 ", err) // Debug
						return
					}

					fmt.Fprintf(w, "Rejection Recorded.")

				}

			} else {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			}

		}

	} else {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}
