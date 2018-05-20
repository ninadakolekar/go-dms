package docs

// ValidateDocName ... Validate Document Name
func ValidateDocName(s string) bool { // If length < 3
	if len(s) <= 2 {
		return false
	}
	return true
}
