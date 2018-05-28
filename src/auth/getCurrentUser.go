package auth

import (
	"fmt"
	"net/http"

	"github.com/ninadakolekar/aizant-dms/src/constants"
)

func GetCurrentUser(r *http.Request) (string, error) {

	// Fetch the cookie
	cookie, err := r.Cookie("session")

	if err != nil {
		fmt.Println("Error Fetching Cookie GetCurrentUser Line 14: ", err) /// Debug
		return "", err
	}

	cookieValue := make(map[string]string)

	err = constants.CookieHandler.Decode("session", cookie.Value, &cookieValue)

	if err != nil {
		fmt.Println("Error decoding cookie GetCurrentUser Line 29: ", err) // Debug
		return "", err
	}

	username := cookieValue["username"]

	return username, nil

}
