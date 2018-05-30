package docs

import (
	"github.com/ninadakolekar/go-dms/src/constants"
	"github.com/ninadakolekar/go-dms/src/models"
	"github.com/ninadakolekar/go-dms/src/utility"
)

// CheckCurrentApprover ... Checks if current user is a approver for a given document
func CheckCurrentApprover(document models.InactiveDoc, username string) bool {
	// If document is not to be reviewed

	if document.FlowStatus != constants.ApproveFlow {
		return false
	}

	// If document has to be reviewed

	// Check if current user is a reviewer

	if !utility.StringInSlice(username, document.Approver) {
		return false
	}

	// Decide according to document type

	if document.DocProcess == "Everyone" || document.DocProcess == "Anyone" {

		hasApproved := utility.StringInSlice(username, document.FlowList)
		return !hasApproved // Return true if has not reviewed already

	} else if document.DocProcess == "OneByOne" {
		if username == document.Approver[int(document.CurrentFlowUser)] {
			return true // Return true if current reviewer is current user
		}
		return false
	}

	return false
}
