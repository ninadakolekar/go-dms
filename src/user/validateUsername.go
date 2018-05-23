package user

import (
	"strings"
)

// ValidateUsername ... Validates username
func ValidateUsername(username string) bool {
	if len(username) <= 2 {
		return false
	}
	if strings.Contains(username, " ") {
		return false
	}
	return true
}
