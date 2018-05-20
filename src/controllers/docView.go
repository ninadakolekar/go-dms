package controllers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ninadakolekar/aizant-dms/src/docs"
)

// DocView ... View mode handler
func DocView(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		vars := mux.Vars(r)
		id := vars["id"]

		if docs.ValidateDocNo(id) {

			tmpl := template.Must(template.ParseFiles("templates/viewDoc.html"))
			tmpl.Execute(w, struct{ DocNumber string }{vars["id"]})

		} else {

			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

		}

	}
}

// func fetchDocBody(id string) ([]string, error) {

// 	document, err := docs.FetchDocByID(id)

// 	if err != nil {
// 		return nil, err
// 	}

// }
