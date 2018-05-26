package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/ninadakolekar/aizant-dms/src/auth"
	user "github.com/ninadakolekar/aizant-dms/src/user"
)

//Dashboard ... dashboard
func Dashboard(w http.ResponseWriter, r *http.Request) {

	usr, err := auth.GetCurrentUser(r)

	fmt.Println(usr) //Debug

	if err != nil { // Login unsucessful
		http.Redirect(w, r, "/", 302)
		return
	}
	UserDetails, err := user.FetchUserByUsername(usr)
	if err != nil { // Login unsucessful
		http.Redirect(w, r, "/", 302)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/dashboard.html"))
	tmpl.Execute(w, UserDetails)
}
