package controllers

import (
	"html/template"
	"net/http"
)

// DocSearch ... Searches for a particulart document
func DocSearch(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/searchDocMultiple.html"))
	tmpl.Execute(w, struct {
		Getb     bool
		Alertb   bool
		Alertmsg string
		Datab    bool
		Data     []string
	}{true, false, "this is alert", false, []string{"no"}})
}
