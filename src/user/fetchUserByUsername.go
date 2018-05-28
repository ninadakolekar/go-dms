package user

import (
	"errors"
	"fmt"

	"github.com/ninadakolekar/aizant-dms/src/constants"
	models "github.com/ninadakolekar/aizant-dms/src/models"
	solr "github.com/rtt/Go-Solr"
)

//FetchUserByUsername ... Fetches the user with given username from the database
func FetchUserByUsername(username string) (models.User, error) {

	var user models.User
	s, err := solr.Init(constants.SolrHost, constants.SolrPort, constants.UserCore)

	if err != nil {
		fmt.Println("ERROR FetchUserByUsername Line 19: ", err) // Debug
		return user, err
	}

	q := solr.Query{
		Params: solr.URLParamMap{
			"q": []string{"id:" + username},
		},
		Rows: 1,
	}

	res, err := s.Select(&q)

	if err != nil {
		fmt.Println("ERROR FetchUserByUsername Line 33: ", err) // Debug
		return user, err
	}

	results := res.Results
	if results.Len() == 0 {
		fmt.Println("ERROR FetchUserByUsername Line 39: User with username " + username + " not found") // Debug
		err := errors.New("User with username " + username + " not found")
		return user, err
	}

	result := results.Get(0)

	if !isUserValid(result) {
		fmt.Println("ERROR FetchUserByUsername Line 46: Invalid Document Mandatory fields found missing") // Debug
		err := errors.New("Invalid User!\nMandatory fields found missing")
		return user, err
	}

	user.Username = result.Field("id").(string)
	user.Name = result.Field("uName").(string)
	user.AvailableInit = result.Field("avInit").(bool)
	user.AvailableCr = result.Field("avCr").(bool)
	user.AvailableApp = result.Field("avAp").(bool)
	user.AvailableAuth = result.Field("avAu").(bool)
	user.AvailableRw = result.Field("avRw").(bool)
	user.AvailableQA = result.Field("avQA").(bool)

	return user, nil
}

func isUserValid(s1 *solr.Document) bool {

	if s1.Field("id") != nil && s1.Field("uName") != nil && s1.Field("avInit") != nil && s1.Field("avCr") != nil && s1.Field("avAp") != nil && s1.Field("avAu") != nil && s1.Field("avRw") != nil && s1.Field("avQA") != nil {
		return true
	}

	return false
}
