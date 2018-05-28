package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	auth "github.com/ninadakolekar/aizant-dms/src/auth"
	"github.com/ninadakolekar/aizant-dms/src/constants"
	docs "github.com/ninadakolekar/aizant-dms/src/docs"
)

// DocCreate ... Handles request to create document
func DocCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		// User Auth Verification

		username, err := auth.GetCurrentUser(r)

		if err != nil { // Auth unsucessful
			fmt.Println("ERROR docView Line 24: ", err) // Debug
			http.Redirect(w, r, "/", 302)
			return
		}

		vars := mux.Vars(r)
		id := vars["id"]

		if docs.ValidateDocNo(id) {

			documentCreator, documentTitle, documentFLow, documentBody, _, err := fetchCreateDocDetails(id)

			// Allow creation only if flow status is 2
			if documentFLow != constants.CreateFlow {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}

			if err != nil {
				fmt.Println("Failed to fetch document: ", err)
				return
			}

			if documentCreator != username { // Not a creator for this document
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}

			tmpl := template.Must(template.ParseFiles("templates/createDoc.html"))
			tmpl.Execute(w, struct {
				DocNumber  string
				DocTitle   string
				DocBody    []string
				DocBodyLen int
			}{id, documentTitle, documentBody, len(documentBody)})

		} else {

			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

		}

	}
}

func fetchCreateDocDetails(id string) (string, string, float64, []string, string, error) {

	document, err := docs.FetchDocByID(id)

	if err != nil {
		return "", "", -1, nil, "", err
	}

	date := strings.Split(strings.Split(document.InitTS, "T")[0], "-")
	initDate := date[2] + "/" + date[1] + "/" + date[0]

	return document.Creator, document.Title, document.FlowStatus, document.DocumentBody, initDate, nil
}
