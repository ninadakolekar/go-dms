package docs

import (
	"github.com/ninadakolekar/aizant-dms/src/constants"
	"github.com/ninadakolekar/aizant-dms/src/models"
	"github.com/ninadakolekar/aizant-dms/src/utility"
)

// CheckCurrentAuthorizer ... Checks if current user is a approver for a given document
func CheckCurrentAuthorizer(document models.InactiveDoc, username string) bool {
	// If document is not to be Authorized

	if document.FlowStatus != constants.AuthFlow {
		return false
	}

	// If document has to be Authorized

	// Check if current user is a Authorizer
	if !utility.StringInSlice(username, document.Authorizer) {
		return false
	}
	// Decide according to document type

	if document.DocProcess == "Everyone" || document.DocProcess == "Anyone" {
		hasAuthorized := utility.StringInSlice(username, document.FlowList)
		return !hasAuthorized // Return true if has not authorized already

	} else if document.DocProcess == "OneByOne" {
		if username == document.Authorizer[int(document.CurrentFlowUser)] {
			return true // Return true if current authorizer is current user
		}
		return false
	}

	return false
}
