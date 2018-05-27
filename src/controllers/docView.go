package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/ninadakolekar/aizant-dms/src/auth"

	"github.com/gorilla/mux"
	"github.com/ninadakolekar/aizant-dms/src/docs"
)

// DocView ... View mode handler
func DocView(w http.ResponseWriter, r *http.Request) {
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

			document, err := docs.FetchDocByID(id)

			if err != nil {
				fmt.Println("Failed to fetch document: ", err)
				return
			}

			EditBtn := false
			if document.Creator == username {
				EditBtn = true
			}

			if document.DocumentBody == nil {
				fmt.Fprintf(w, "<script>alert('Document not created!');</script>")
			} else {
				tmpl := template.Must(template.ParseFiles("templates/viewDoc.html"))
				tmpl.Execute(w, struct {
					DocNumber   string
					DocTitle    string
					DocBody     []string
					DocInitDate string

					// Extra Details
					DocDept    string
					DocType    string
					DocEffDate string
					DocExpDate string
					Edit       bool
					Rw         bool
				}{id, document.Title, document.DocumentBody, document.InitTS, document.DocDept, document.DocType, document.DocEffDate, document.DocExpDate, EditBtn, docs.CheckCurrentReviewer(document, username)})
			}

		} else {

			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

		}

	}
}
