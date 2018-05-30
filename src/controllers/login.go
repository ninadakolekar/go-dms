package controllers

import (
	"html/template"
	"net/http"

	auth "github.com/ninadakolekar/go-dms/src/auth"
)

//Login ... login
func Login(w http.ResponseWriter, r *http.Request) {

	// User Auth Verification

	_, err := auth.GetCurrentUser(r)

	if err == nil { // Already logged-in
		http.Redirect(w, r, "/dashboard", 302)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/login.html"))
	tmpl.Execute(w, nil)
}
