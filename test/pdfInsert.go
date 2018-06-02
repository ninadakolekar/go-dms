package test

import (
	"strconv"

	docs "github.com/ninadakolekar/go-dms/src/docs"
	"github.com/ninadakolekar/go-dms/src/models"
	utility "github.com/ninadakolekar/go-dms/src/utility"
)

func InsertPDF(papers []paper) {

	authorsList := getAuthors(papers)

	for _, p := range papers {

		go func(p paper) {

			var document models.InactiveDoc

			document.DocNo = p.id
			document.Title = p.title
			document.Initiator, document.Creator, document.Reviewer, document.Approver, document.Authorizer = selectAuthors(p, authorsList)
			document.DocDept = p.venue
			document.DocEffDate = makeDate(p.publDate, 0)
			document.DocExpDate = makeDate(p.publDate, 4)
			document.InitTS = makeDate(p.publDate, -1)

			document.CurrentFlowUser = 0
			document.FlowList = nil
			document.FlowStatus = float64(utility.RandomInteger(2, 5))

			document.DocProcess = getRandomDocProcess()
			document.DocStatus = false

			document.DocType = getRandomDocType()

			document.DocumentBody, _ = ConvertURL2Strings(p.url, p.title)

			document.QA = "firefox"

			docs.AddInactiveDoc(document)

		}(p)
	}

}

func selectAuthors(p paper, authors []string) (string, string, []string, []string, []string) {

	initiator := getUsername(p.authors[0])
	creator := getUsername(p.authors[1])
	reviewer := []string{getUsername(p.authors[2])}
	app := []string{}
	auth := []string{}

	if len(p.authors) == 3 {

		app = getRandomUsers(authors, 3)
		auth = nil

	} else if len(p.authors) == 4 {

		app = append(app, getUsername(p.authors[3]))
		auth = nil

	} else if len(p.authors) >= 5 {

		auth = append(auth, getUsername(p.authors[3]))

		for i := 4; i < len(p.authors); i++ {
			app = append(app, getUsername(p.authors[i]))
		}

	}

	return initiator, creator, reviewer, app, auth

}

func getRandomUsers(authors []string, maxUsers int) []string {

	maxUsers = utility.RandomInteger(1, maxUsers)
	users := []string{}
	for i := 0; i < maxUsers; i++ {
		j := utility.RandomInteger(0, len(authors))
		users = append(users, getUsername(authors[j]))
	}
	return users
}

func getRandomDocProcess() string {
	processID := utility.RandomInteger(1, 3)
	if processID == 1 {
		return "Everyone"
	} else if processID == 2 {
		return "Anyone"
	} else if processID == 3 {
		return "OneByOne"
	}

	return ""
}

func makeDate(date string, offset int) string {
	year, _ := strconv.Atoi(date)
	year = year + offset
	date = strconv.Itoa(year) + "-" + strconv.Itoa(utility.RandomInteger(10, 12)) + "-" + strconv.Itoa(utility.RandomInteger(10, 30))
	return utility.XMLDate(date)
}

func getRandomDocType() string {
	processID := utility.RandomInteger(1, 3)
	if processID == 1 {
		return "SOP"
	} else if processID == 2 {
		return "HR"
	} else if processID == 3 {
		return "STP"
	}

	return ""
}
