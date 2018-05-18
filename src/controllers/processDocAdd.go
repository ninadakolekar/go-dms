package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	constant "github.com/ninadakolekar/aizant-dms/src/constants"
	doc "github.com/ninadakolekar/aizant-dms/src/docs"
	model "github.com/ninadakolekar/aizant-dms/src/models"
	solr "github.com/rtt/Go-Solr"
)

// ProcessDocAdd ... Process the form-values and add the document
func ProcessDocAdd(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		initTime := xmlTimeNow()

		//TODO : Sanitize the form data

		r.ParseForm()

		docNumber := r.Form["docNumber"][0]
		docName := r.Form["docName"][0]
		docProcess := r.Form["docProcess"][0]
		docType := r.Form["docType"][0]
		docDept := r.Form["docDept"][0]
		docEffDate := r.Form["docEffDate"][0]
		docExpDate := r.Form["docExpDate"][0]
		docCreator := r.Form["docCreator"][0]
		docAuth := r.Form["docAuth"]
		docReviewers := r.Form["docReviewers"]
		docApprovers := r.Form["docApprovers"]

		fmt.Println("Form Received\n ", docNumber, docName, docProcess, docType, docDept, xmlDate(docEffDate), xmlDate(docExpDate), docCreator, docReviewers, docApprovers, docAuth, initTime) // Debug

		// Server-side validation

		if validateDocNo(docNumber) && validateDocName(docName) {
			// Make a new inactiveDoc struct using received form data

			newDoc := model.InactiveDoc{
				DocNo:        docNumber,
				Title:        docName,
				DocType:      docType,
				DocProcess:   docProcess,
				DocEffDate:   xmlDate(docEffDate),
				DocExpDate:   xmlDate(docExpDate),
				DocStatus:    false,
				Initiator:    "self", // Initiator is "self" currently
				Creator:      docCreator,
				Reviewer:     docReviewers,
				Approver:     docApprovers,
				Authorizer:   docAuth,
				DocDept:      docDept,
				FlowStatus:   0,
				DocTemplate:  0,
				InitTS:       initTime,
				DocumentBody: "Empty Body",
			}
			// Insert the new document
			resp, err := doc.AddInactiveDoc(newDoc)

			// Respond
			if err != nil {
				fmt.Println("ERROR ProcessDocAdd() Line 47: " + err.Error()) // Debug
				fmt.Fprintf(w, "<script>alert('Failed to create new document.');</script>")
			} else {
				fmt.Println(resp) // Debug
				fmt.Fprintf(w, "<script>alert('New document successfully created.');</script>")
			}

		} else {
			fmt.Fprintf(w, "<script>alert('Failed to create new document (ERROR: Invalid Document Number or Name).');</script>")
		}
	}

	// Render a new form
	tmpl := template.Must(template.ParseFiles("templates/addNewDoc.html"))

	tmpl.Execute(w, templateData{Approvers: SendApprovers(), Reviewers: SendReviewers(), Authorisers: SendAuthoriser(), Creators: SendCreators()})
}

// Validate Document Number
func validateDocNo(str string) bool {
	if len(str) <= 2 || strings.Contains(str, " ") { // If length < 3 or if contains whitespace
		return false
	}
	s, err := solr.Init(constant.SolrHost, constant.SolrPort, constant.DocsCore)

	if err != nil {
		fmt.Println(err)
		return false
	}
	quer := "id:" + str
	q := solr.Query{ //checking in backend whether any other documnet with same id is present

		Params: solr.URLParamMap{
			"q": []string{quer},
		},
		Rows: 1,
	}
	res, err := s.Select(&q)
	if err != nil {
		fmt.Println(err)
	}

	results := res.Results
	if results.Len() != 0 {
		return false
	}
	return true
}

// Validate Document Name

func validateDocName(s string) bool { // If length < 3
	if len(s) <= 2 {
		return false
	}
	return true
}

// Get current time in XML format
func xmlTimeNow() string {
	t := time.Now().UTC().String()
	x := strings.Fields(t)
	xmlTime := x[0] + "T" + x[1] + "Z"
	return xmlTime
}

// Get Date in XML format
func xmlDate(date string) string {
	return date + "T23:59:59"
}
