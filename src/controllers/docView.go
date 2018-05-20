package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/ninadakolekar/aizant-dms/src/docs"
)

// DocView ... View mode handler
func DocView(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		vars := mux.Vars(r)
		id := vars["id"]

		if docs.ValidateDocNo(id) {

			documentTitle, documentBody, documentInitDate, err := fetchViewDocDetails(id)

			if err != nil {
				fmt.Println("Failed to fetch document: ", err)
				return
			}

			if documentBody == nil {
				fmt.Fprintf(w, "<script>alert('Document not created!');</script>")
			} else {
				tmpl := template.Must(template.ParseFiles("templates/viewDoc.html"))
				tmpl.Execute(w, struct {
					DocNumber   string
					DocTitle    string
					DocBody     []string
					DocInitDate string
				}{id, documentTitle, documentBody, documentInitDate})
			}

		} else {

			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

		}

	}
}

func fetchViewDocDetails(id string) (string, []string, string, error) {

	document, err := docs.FetchDocByID(id)

	if err != nil {
		return "", nil, "", err
	}

	date := strings.Split(strings.Split(document.InitTS, "T")[0], "-")
	initDate := date[2] + "/" + date[1] + "/" + date[0]

	return document.Title, document.DocumentBody, initDate, nil
}
