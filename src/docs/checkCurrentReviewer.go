package docs

import (
	"github.com/ninadakolekar/go-dms/src/constants"
	models "github.com/ninadakolekar/go-dms/src/models"
	"github.com/ninadakolekar/go-dms/src/utility"
)

// CheckCurrentReviewer ... Checks if current user is a reviewer for a given document
func CheckCurrentReviewer(document models.InactiveDoc, username string) bool {

	// If document is not to be reviewed

	if document.FlowStatus != constants.ReviewFlow {
		return false
	}

	// If document has to be reviewed

	// Check if current user is a reviewer

	if !utility.StringInSlice(username, document.Reviewer) {
		return false
	}

	// Decide according to document type

	if document.DocProcess == "Everyone" || document.DocProcess == "Anyone" {

		hasReviewed := utility.StringInSlice(username, document.FlowList)
		return !hasReviewed // Return true if has not reviewed already

	} else if document.DocProcess == "OneByOne" {
		if username == document.Reviewer[int(document.CurrentFlowUser)] {
			return true // Return true if current reviewer is current user
		}
		return false
	}

	return false
}
