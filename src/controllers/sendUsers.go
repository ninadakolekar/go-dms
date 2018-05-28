package controllers

import (
	"fmt"

	"github.com/ninadakolekar/aizant-dms/src/constants"
	solr "github.com/rtt/Go-Solr"
)

type Strings2 struct {
	Name string
	Id   string
}

//SendApprovers ... Returns the list of available Approvers
func SendApprovers() []Strings2 {
	str := []Strings2{}

	s, err := solr.Init(constants.SolrHost, constants.SolrPort, constants.UserCore)
	if err != nil {
		fmt.Println(err)
		return str
	}

	q := solr.Query{ //checking in backend whether any other documnet with same id is present

		Params: solr.URLParamMap{
			"q": []string{"avAp:true"},
		},
		Rows: 10,
	}

	res, err := s.Select(&q)
	if err != nil {
		fmt.Println(err)
	}

	results := res.Results

	if results.Len() > 0 {
		len := results.Len()
		for i := 0; i < len; i++ {
			str = append(str, Strings2{Name: results.Get(i).Field("uName").(string), Id: results.Get(i).Field("id").(string)})
		}
	}

	return str
}

// SendReviewers ... Returns the list of available Reviewers
func SendReviewers() []Strings2 {
	str := []Strings2{}

	s, err := solr.Init(constants.SolrHost, constants.SolrPort, constants.UserCore)
	if err != nil {
		fmt.Println(err)
		return str
	}

	q := solr.Query{

		Params: solr.URLParamMap{
			"q": []string{"avRw:true"},
		},
		Rows: 10,
	}

	res, err := s.Select(&q)
	if err != nil {
		fmt.Println(err)
	}

	results := res.Results

	if results.Len() > 0 {
		len := results.Len()
		for i := 0; i < len; i++ {
			str = append(str, Strings2{Name: results.Get(i).Field("uName").(string), Id: results.Get(i).Field("id").(string)})
		}
	}

	return str
}

// SendCreators ... Returns the list of available Creators
func SendCreators() []Strings2 {
	str := []Strings2{}

	s, err := solr.Init(constants.SolrHost, constants.SolrPort, constants.UserCore)
	if err != nil {
		fmt.Println(err)
		return str
	}

	q := solr.Query{ //checking in backend whether any other documnet with same id is present

		Params: solr.URLParamMap{
			"q": []string{"avCr:true"},
		},
		Rows: 10,
	}

	res, err := s.Select(&q)
	if err != nil {
		fmt.Println(err)
	}

	results := res.Results

	if results.Len() > 0 {
		len := results.Len()
		for i := 0; i < len; i++ {
			str = append(str, Strings2{Name: results.Get(i).Field("uName").(string), Id: results.Get(i).Field("id").(string)})
		}
	}

	return str
}

// SendAuthoriser ... Returns the list of available authorizers
func SendAuthoriser() []Strings2 {
	str := []Strings2{}

	s, err := solr.Init(constants.SolrHost, constants.SolrPort, constants.UserCore)
	if err != nil {
		fmt.Println(err)
		return str
	}

	q := solr.Query{ //checking in backend whether any other documnet with same id is present

		Params: solr.URLParamMap{
			"q": []string{"avAu:true"},
		},
		Rows: 10,
	}

	res, err := s.Select(&q)
	if err != nil {
		fmt.Println(err)
	}

	results := res.Results

	if results.Len() > 0 {
		len := results.Len()
		for i := 0; i < len; i++ {
			str = append(str, Strings2{Name: results.Get(i).Field("uName").(string), Id: results.Get(i).Field("id").(string)})
		}
	}

	return str
}
