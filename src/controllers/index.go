package controllers

import (
	"fmt"
	"net/http"
)

// Index ... Handles index view
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)

}
