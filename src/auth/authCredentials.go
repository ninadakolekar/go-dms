package auth

import (
	"errors"
	"fmt"

	"github.com/ninadakolekar/aizant-dms/src/constants"
	solr "github.com/rtt/Go-Solr"
)

//AuthCredentials ... authunticates users
func AuthCredentials(username string, password string) (bool, error) {

	s, err := solr.Init(constants.SolrHost, constants.SolrPort, constants.UserCore)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	q := solr.Query{
		Params: solr.URLParamMap{
			"q": []string{"id:" + username},
		},
		Rows: 1,
	}
	res, err := s.Select(&q)

	if err != nil {
		fmt.Println(err)
		return false, err
	}
	results := res.Results
	if results.Len() == 0 {
		err := errors.New("Document with id " + username + "not found")
		return false, err
	}
	if validatePassword(password) == false {
		err := errors.New("password incorrect")
		return false, err
	}
	return true, nil
}
