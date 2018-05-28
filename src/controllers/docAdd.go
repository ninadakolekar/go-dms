package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	auth "github.com/ninadakolekar/aizant-dms/src/auth"
	"github.com/ninadakolekar/aizant-dms/src/models"
	user "github.com/ninadakolekar/aizant-dms/src/user"
)

//DocAdd ...
func DocAdd(w http.ResponseWriter, r *http.Request) {

	// User Auth Verification

	username, err := auth.GetCurrentUser(r)

	if err != nil { // Auth unsucessful
		fmt.Println("ERROR DocAdd Line 23: ", err) // Debug
		http.Redirect(w, r, "/", 302)
		return
	}

	user, err := user.FetchUserByUsername(username)

	if err != nil { // User fetch unsucessful
		fmt.Println("ERROR DocAdd Line 31: ", err) // Debug
		http.Redirect(w, r, "/", 302)
		return
	}

	if user.AvailableInit == false {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/addNewDoc.html"))
	tmpl.Execute(w, docAddMsg{Username: username, Datab: false, Errb: false, Datamsg: "hi", Approvers: SendApprovers(), Reviewers: SendReviewers(), Authorisers: SendAuthoriser(), Creators: SendCreators(), DocumentExist: false, Redirect: false, Document: models.InactiveDoc{}})
}

type docAddMsg struct {
	Username      string
	Datab         bool
	Errb          bool
	Datamsg       string
	Approvers     []Strings2
	Reviewers     []Strings2
	Authorisers   []Strings2
	Creators      []Strings2
	DocumentExist bool
	Redirect      bool
	Document      models.InactiveDoc
}
