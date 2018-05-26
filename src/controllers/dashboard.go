package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/ninadakolekar/aizant-dms/src/user"
)

//Login ... login
func Dashboard(w http.ResponseWriter, r *http.Request) {

	user, err := user.GetCurrentUser(r)

	fmt.Println(user) //Debug

	if err != nil { // Login unsucessful
		http.Redirect(w, r, "/", 302)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/dashboard.html"))
	tmpl.Execute(w, struct {
		Username string
	}{user})
}
