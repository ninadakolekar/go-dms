package controllers

import (
	"html/template"
	"net/http"

	constant "github.com/ninadakolekar/aizant-dms/src/constants"
)

//DocAdd ... not competed
func DocAdd(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/addDoc.html"))
	tmpl.Execute(w, struct{ s uint }{constant.MinDocNumLen})
}
