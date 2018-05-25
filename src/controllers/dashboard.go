package controllers

import (
	"html/template"
	"net/http"
)

//Login ... login
func Dashboard(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/dashboard.html"))
	tmpl.Execute(w, nil)
}
