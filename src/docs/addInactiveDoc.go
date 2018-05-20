package docs

import (
	"fmt" // Debug

	constant "github.com/ninadakolekar/aizant-dms/src/constants"
	model "github.com/ninadakolekar/aizant-dms/src/models"
	solr "github.com/rtt/Go-Solr"
)

// AddInactiveDoc ... Adds a new document to Inactive Docs
func AddInactiveDoc(doc model.InactiveDoc) (*solr.UpdateResponse, error) {

	// Initialize a solr connection
	s, err := solr.Init(constant.SolrHost, constant.SolrPort, constant.DocsCore)

	if err != nil { // If connection fails
		fmt.Println("ERROR AddInactiveDoc Line 14: ", err) // Debug
		return nil, err
	}

	f := map[string]interface{}{
		"add": []interface{}{
			map[string]interface{}{
				"id":            doc.DocNo,
				"title":         doc.Title,
				"docProcess":    doc.DocProcess,
				"docStatus":     doc.DocStatus,
				"approver":      doc.Approver,
				"authorizer":    doc.Authorizer,
				"creator":       doc.Creator,
				"initiator":     doc.Initiator,
				"docDepartment": doc.DocDept,
				"docTemplateID": doc.DocTemplate,
				"body":          doc.DocumentBody,
				"docType":       doc.DocType,
				"flowStatus":    doc.FlowStatus,
				"reviewer":      doc.Reviewer,
				"initTime":      doc.InitTS,
				"createTime":    doc.CreateTS,
				"reviewTime":    doc.ReviewTS,
				"approveTime":   doc.ApproveTS,
				"authTime":      doc.AuthTS,
				"effDate":       doc.DocEffDate,
				"expDate":       doc.DocExpDate,
			},
		},
	}

	// Send off the update to solr (2nd parameter indicates we also want to commit the operation)
	resp, err := s.Update(f, true)

	if err != nil {
		fmt.Println("ERROR addInactiveDoc Line 46: ", err) // Debug
		return resp, err
	}

	return resp, nil

}
