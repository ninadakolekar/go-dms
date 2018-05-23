package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/securecookie"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

// ProcessLogin ... Processes Login form

func ProcessLogin(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	password := r.FormValue("password")

	redirectTarget := "/"

	if validateCredentials(username, password) {

		// Check Credentials

		setSession(username, w)
		redirectTarget = "/dashboard"

	}

	http.Redirect(w, r, redirectTarget, 302)
}

func validateCredentials(username string, password string) bool {

	if username == "" || password == "" {
		return false
	}

	return true
}

func setSession(username string, response http.ResponseWriter) {

	value := map[string]string{
		"name": username,
	}

	encoded, err := cookieHandler.Encode("session", value)

	if err != nil {
		fmt.Println("Error setSession Line 36: ", err)
		return
	}

	cookie := &http.Cookie{
		Name:  "session",
		Value: encoded,
		Path:  "/",
	}

	http.SetCookie(response, cookie)
}
