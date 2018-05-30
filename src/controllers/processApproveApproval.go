package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	auth "github.com/ninadakolekar/go-dms/src/auth"
	"github.com/ninadakolekar/go-dms/src/constants"
	"github.com/ninadakolekar/go-dms/src/docs"
)

// ProcessApproveApproval ... Handles Approver's Approval/Rejection
func ProcessApproveApproval(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		// User Auth Verification

		username, err := auth.GetCurrentUser(r)

		if err != nil { // Auth unsucessful
			fmt.Println("ERROR ProcessApproveApproval Line 23: ", err) // Debug
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

			if !docs.CheckCurrentApprover(Document, username) || Document.FlowStatus != constants.ApproveFlow {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}

			apResponse := r.FormValue("ap-answer")

			if apResponse == "approve" {

				if Document.DocProcess == constants.Everyone {

					Document.FlowList = append(Document.FlowList, username)

					if len(Document.FlowList) == len(Document.Approver) {
						if Document.Authorizer == nil {
							Document.FlowStatus = constants.ActiveFlow
							Document.DocStatus = true
						} else {
							Document.FlowStatus = constants.AuthFlow
						}
						Document.FlowList = nil
						Document.CurrentFlowUser = 0
					}

					_, err := docs.AddInactiveDoc(Document)

					if err != nil {
						fmt.Println("Error: Document Update Failed ProcessApproveApproval Line 67 ", err) // Debug
						return
					}

					fmt.Fprintf(w, "Approval Recorded.")

				} else if Document.DocProcess == constants.Anyone {

					if Document.Authorizer == nil {
						Document.FlowStatus = constants.ActiveFlow
						Document.DocStatus = true
					} else {
						Document.FlowStatus = constants.AuthFlow
					}
					Document.FlowList = nil
					Document.CurrentFlowUser = 0

					_, err := docs.AddInactiveDoc(Document)

					if err != nil {
						fmt.Println("Error: Document Update Failed ProcessApproveApproval Line 87 ", err) // Debug
						return
					}

					fmt.Fprintf(w, "Approval Recorded.")

				} else if Document.DocProcess == constants.OneByOne {

					Document.CurrentFlowUser++

					if int(Document.CurrentFlowUser) == len(Document.Approver) {
						if Document.Authorizer == nil {
							Document.FlowStatus = constants.ActiveFlow
							Document.DocStatus = true
						} else {
							Document.FlowStatus = constants.AuthFlow
						}
						Document.FlowList = nil
						Document.CurrentFlowUser = 0
					}

					_, err := docs.AddInactiveDoc(Document)

					if err != nil {
						fmt.Println("Error: Document Update Failed ProcessApproveApproval Line 111 ", err) // Debug
						return
					}

					fmt.Fprintf(w, "Approval Recorded.")

				}

			} else if apResponse == "reject" {

				if Document.DocProcess == constants.Everyone {

					Document.FlowStatus = constants.CreateFlow
					Document.FlowList = nil
					Document.CurrentFlowUser = 0

					_, err := docs.AddInactiveDoc(Document)

					if err != nil {
						fmt.Println("Error: Document Update Failed ProcessApproveApproval Line 130 ", err) // Debug
						return
					}

					fmt.Fprintf(w, "Rejection Recorded.")

				} else if Document.DocProcess == constants.Anyone {

					Document.FlowList = append(Document.FlowList, username)

					if len(Document.FlowList) == len(Document.Approver) {
						Document.FlowStatus = constants.CreateFlow
						Document.FlowList = nil
						Document.CurrentFlowUser = 0
					}

					_, err := docs.AddInactiveDoc(Document)

					if err != nil {
						fmt.Println("Error: Document Update Failed ProcessApproveApproval Line 149 ", err) // Debug
						return
					}

					fmt.Fprintf(w, "Rejection Recorded.")

				} else if Document.DocProcess == constants.OneByOne {

					Document.FlowStatus = constants.CreateFlow
					Document.FlowList = nil
					Document.CurrentFlowUser = 0

					_, err := docs.AddInactiveDoc(Document)

					if err != nil {
						fmt.Println("Error: Document Update Failed ProcessApproveApproval Line 164 ", err) // Debug
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
