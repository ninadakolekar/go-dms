package controllers

import (
	"html/template"
	"net/http"
)

//Login ... login
func Login(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/login.html"))
	tmpl.Execute(w, nil)
}
