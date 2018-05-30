package controllers

import (
	"fmt"
	"net/http"

	auth "github.com/ninadakolekar/go-dms/src/auth"
)

// ProcessLogout ... Handles logout request
func ProcessLogout(w http.ResponseWriter, r *http.Request) {

	// User Auth Verification

	_, err := auth.GetCurrentUser(r)

	if err != nil { // Auth unsucessful
		fmt.Println("ERROR ProcessLogout Line 18: ", err) // Debug
		http.Redirect(w, r, "/", 302)
		return
	}

	clearSession(w)
	http.Redirect(w, r, "/", 302)
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}
