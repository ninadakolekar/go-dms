package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	auth "github.com/ninadakolekar/go-dms/src/auth"
	"github.com/ninadakolekar/go-dms/src/docs"
	user "github.com/ninadakolekar/go-dms/src/user"
)

// DocAddEdit ... Re-initiating a document
func DocAddEdit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		// User Auth Verification

		username, err := auth.GetCurrentUser(r)

		if err != nil { // Auth unsucessful
			fmt.Println("ERROR DocAddEdit Line 23: ", err) // Debug
			http.Redirect(w, r, "/", 302)
			return
		}

		user, err := user.FetchUserByUsername(username)

		if err != nil { // User fetch unsucessful
			fmt.Println("ERROR DocAddEdit Line 31: ", err) // Debug
			http.Redirect(w, r, "/", 302)
			return
		}

		vars := mux.Vars(r)
		id := vars["id"]

		if docs.ValidateDocNo(id) {

			document, err := docs.FetchDocByID(id)

			if err != nil {
				fmt.Println("Failed to fetch document: ", err)
				return
			}

			if user.AvailableInit == false || username != document.Initiator {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}

			tmpl := template.Must(template.ParseFiles("templates/addNewDoc.html"))
			tmpl.Execute(w, docAddMsg{Datab: false, Errb: false, Datamsg: "hi", Approvers: SendApprovers(), Reviewers: SendReviewers(), Authorisers: SendAuthoriser(), Creators: SendCreators(), DocumentExist: true, Redirect: false, Document: document})

		} else {

			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

		}

	}
}
