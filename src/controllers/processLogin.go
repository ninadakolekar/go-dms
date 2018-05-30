package controllers

import (
	"fmt"
	"log"
	"net/http"

	auth "github.com/ninadakolekar/go-dms/src/auth"
	"github.com/ninadakolekar/go-dms/src/constants"
)

// ProcessLogin ... Processes Login form
func ProcessLogin(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		username := r.FormValue("username")
		password := r.FormValue("password")
		redirectTarget := "/"

		if auth.ValidateLoginCredentials(username, password) {

			_, err := auth.AuthCredentials(username, password)
			if err == nil {
				setSession(username, w)
				redirectTarget = "/dashboard"
			} else {
				log.Println("incorrect username or password")
			}
		}

		http.Redirect(w, r, redirectTarget, 302)
	}
}

func setSession(username string, response http.ResponseWriter) {

	value := map[string]string{
		"username": username,
	}

	encoded, err := constants.CookieHandler.Encode("session", value)

	if err != nil {
		fmt.Println("Error setSession Line 36: ", err) // Debug
		return
	}

	cookie := &http.Cookie{
		Name:  "session",
		Value: encoded,
		Path:  "/",
	}

	http.SetCookie(response, cookie)
}
