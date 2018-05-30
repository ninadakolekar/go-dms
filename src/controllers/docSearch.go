package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	auth "github.com/ninadakolekar/go-dms/src/auth"
)

// DocSearch ... Searches for a particulart document
func DocSearch(w http.ResponseWriter, r *http.Request) {

	// User Auth Verification

	_, err := auth.GetCurrentUser(r)

	if err != nil { // Auth unsucessful
		fmt.Println("ERROR docView Line 24: ", err) // Debug
		http.Redirect(w, r, "/", 302)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/newSearch.html"))
	tmpl.Execute(w, struct {
		Getb     bool
		Alertb   bool
		Alertmsg string
		Datab    bool
		Data     []string
	}{true, false, "this is alert", false, []string{"no"}})
}
