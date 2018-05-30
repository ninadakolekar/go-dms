package test

import (
	doc "github.com/ninadakolekar/go-dms/src/docs"
)

// Deletes multiple documents at once
func DeleteDocs(docIDs []string) {
	for _, docID := range docIDs {
		doc.DeleteInactiveDoc(docID)
	}
}
