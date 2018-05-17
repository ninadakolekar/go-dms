package controllers

import (
	"html/template"
	"net/http"
)

// DocSearch ... Searches for a particulart document
func DocSearch(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/searchDoc.html"))
	tmpl.Execute(w, struct{ s bool }{true})
}
