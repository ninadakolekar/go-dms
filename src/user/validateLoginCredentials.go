package user

// ValidateLoginCredentials ... Validates username and password
func ValidateLoginCredentials(username string, password string) bool {
	return ValidateUsername(username) && validatePassword(password)
}
