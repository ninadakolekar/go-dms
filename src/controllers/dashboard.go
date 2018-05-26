package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/ninadakolekar/aizant-dms/src/auth"
	"github.com/ninadakolekar/aizant-dms/src/models"
)

//Dashboard ... dashboard
func Dashboard(w http.ResponseWriter, r *http.Request) {

	user, err := auth.GetCurrentUser(r)

	fmt.Println(user) //Debug

	if err != nil { // Login unsucessful
		http.Redirect(w, r, "/", 302)
		return
	}

	UserDetails := models.User{Username: user, Name: "Loremipsum", AvailableApp: true, AvailableAuth: false, AvailableCr: true, AvailableInit: false, AvailableQA: false, AvailableRvw: true}

	tmpl := template.Must(template.ParseFiles("templates/dashboard.html"))
	tmpl.Execute(w, UserDetails)
}
