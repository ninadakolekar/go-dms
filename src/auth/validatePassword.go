package auth

// validatePassword ... Validates password
func validatePassword(password string) bool {
	if len(password) <= 2 {
		return false
	}
	return true
}
