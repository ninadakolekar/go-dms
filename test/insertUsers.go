package test

import (
	"strings"

	"github.com/ninadakolekar/go-dms/src/models"
	"github.com/ninadakolekar/go-dms/src/user"
)

func InsertUsers(papers []paper) {

	authors := getAuthors(papers)

	for _, author := range authors {
		go func(author string) {

			user.AddUser(getUserModel(author))

		}(author)
	}
}

func getUserModel(author string) models.User {

	var newUser models.User

	newUser.Name = author
	newUser.Username = getUsername(author)

	newUser.AvailableInit = true
	newUser.AvailableCr = true
	newUser.AvailableRw = true
	newUser.AvailableApp = true
	newUser.AvailableAuth = true
	newUser.AvailableQA = false

	return newUser
}

func getUsername(author string) string {

	return strings.Replace(author, " ", "_", -1)
}
