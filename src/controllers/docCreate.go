package controllers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	docs "github.com/ninadakolekar/aizant-dms/src/docs"
)

// DocCreate ... Handles request to create document
func DocCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		vars := mux.Vars(r)
		id := vars["id"]

		if docs.ValidateDocNo(id) {

			tmpl := template.Must(template.ParseFiles("templates/createDoc.html"))
			tmpl.Execute(w, struct{ DocNumber string }{vars["id"]})

		} else {

			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

		}

	}
}
