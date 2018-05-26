package test

import "github.com/ninadakolekar/aizant-dms/src/user"

// DeleteAllUsers ... Delete multiple users at once
func DeleteAllUsers(usernames []string) {
	for _, username := range usernames {
		user.DeleteUser(username)
	}
}
