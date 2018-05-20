package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	docs "github.com/ninadakolekar/aizant-dms/src/docs"
)

// DocCreate ... Handles request to create document
func DocCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		vars := mux.Vars(r)
		id := vars["id"]

		if docs.ValidateDocNo(id) {

			documentTitle, _, _, err := fetchCreateDocDetails(id)

			if err != nil {
				fmt.Println("Failed to fetch document: ", err)
				return
			}

			tmpl := template.Must(template.ParseFiles("templates/createDoc.html"))
			tmpl.Execute(w, struct {
				DocNumber string
				DocTitle  string
			}{id, documentTitle})

		} else {

			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

		}

	}
}

func fetchCreateDocDetails(id string) (string, []string, string, error) {

	document, err := docs.FetchDocByID(id)

	if err != nil {
		return "", nil, "", err
	}

	date := strings.Split(strings.Split(document.InitTS, "T")[0], "-")
	initDate := date[2] + "/" + date[1] + "/" + date[0]

	return document.Title, document.DocumentBody, initDate, nil
}
