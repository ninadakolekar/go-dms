package controllers

import (
	"fmt"
	"html"
	"net/http"
	"strconv"

	auth "github.com/ninadakolekar/aizant-dms/src/auth"
	"github.com/ninadakolekar/aizant-dms/src/docs"
	utility "github.com/ninadakolekar/aizant-dms/src/utility"
)

// ProcessDocCreate ... Processes Document Create Form Input
func ProcessDocCreate(w http.ResponseWriter, r *http.Request) {

	// User Auth Verification

	username, err := auth.GetCurrentUser(r)

	if err != nil { // Auth unsucessful
		fmt.Println("ERROR ProcessDocCreate Line 22: ", err) // Debug
		http.Redirect(w, r, "/", 302)
		return
	}

	if r.Method == "POST" {

		createTime := utility.XMLTimeNow()

		docNumber := html.EscapeString(r.FormValue("docNumber"))

		// Server-side validation of doc number

		if docs.ValidateDocNo(docNumber) {

			paraCount, err := strconv.Atoi(html.EscapeString(r.FormValue("paraCount")))
			if err != nil {
				fmt.Println("Invalid Paragraph Count!", err)
				fmt.Println("Invalid Paragraph Count! (paracount) ProcessDocCreate Line 25", err) // Debug
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			}

			documentBody := make([]string, paraCount)

			for i := 1; i <= paraCount; i++ {
				currentName := "para" + strconv.Itoa(i)
				documentBody[i-1] = html.EscapeString(r.FormValue(currentName))
			}

			document, err := docs.FetchDocByID(docNumber)

			if document.Creator != username { // Not a creator for this document
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}

			if err != nil {
				fmt.Println("ERROR Fetching document ProcessDocCreate Line 40") // Debug
				fmt.Println(err)
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			}

			document.DocumentBody = documentBody
			document.CreateTS = createTime

			resp, err := docs.AddInactiveDoc(document)

			// Respond
			if err != nil {
				fmt.Println("ERROR ProcessDocAdd() Line 57: " + err.Error()) // Debug
				fmt.Fprintf(w, "<script>alert('Failed to create document.');window.location.replace('/dashboard');</script>")
				http.Redirect(w, r, "/dashboard", 301)
			} else {
				fmt.Println(resp) // Debug
				fmt.Fprintf(w, "<script>alert('Document successfully created.');window.location.replace('/dashboard');</script>")
				http.Redirect(w, r, "/dashboard", 301)
			}
		} else {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}

	}

}
