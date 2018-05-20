package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ninadakolekar/aizant-dms/src/docs"
	utility "github.com/ninadakolekar/aizant-dms/src/utility"
)

// ProcessDocCreate ... Processes Document Create Form Input
func ProcessDocCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		createTime := utility.XMLTimeNow()

		docNumber := r.FormValue("docNumber")

		// Server-side validation of doc number

		if docs.ValidateDocNo(docNumber) {

			paraCount, err := strconv.Atoi(r.FormValue("paraCount"))
			if err != nil {
				fmt.Println("Invalid Paragraph Count!", err)
				fmt.Println("Invalid Paragraph Count! (paracount) ProcessDocCreate Line 25", err) // Debug
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			}

			documentBody := make([]string, paraCount)

			for i := 1; i <= paraCount; i++ {
				currentName := "para" + strconv.Itoa(i)
				documentBody[i-1] = r.FormValue(currentName)
			}

			fmt.Println(docNumber, documentBody) // Debug

			document, err := docs.FetchDocByID(docNumber)
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
				fmt.Fprintf(w, "<script>alert('Failed to create document.');</script>")
			} else {
				fmt.Println(resp) // Debug
				fmt.Fprintf(w, "<script>alert('Document successfully created.');</script>")
			}
		} else {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}

	}

}
