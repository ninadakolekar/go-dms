package controllers

import (
	"fmt"

	solr "github.com/rtt/Go-Solr"
)

//SendApprovers ... sends present approvers in users
func SendApprovers() []string {
	str := []string{}

	s, err := solr.Init("localhost", 8983, "user")
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
			str = append(str, results.Get(i).Field("uName").(string))
		}
	}

	return str
}

//SendReviewers ... sends reviewers
func SendReviewers() []string {
	str := []string{}

	s, err := solr.Init("localhost", 8983, "user")
	if err != nil {
		fmt.Println(err)
		return str
	}

	q := solr.Query{ //checking in backend whether any other documnet with same id is present

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
			str = append(str, results.Get(i).Field("uName").(string))
		}
	}

	return str
}
func SendCreators() []string {
	str := []string{}

	s, err := solr.Init("localhost", 8983, "user")
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
			str = append(str, results.Get(i).Field("uName").(string))
		}
	}

	return str
}
func SendAuthoriser() []string {
	str := []string{}

	s, err := solr.Init("localhost", 8983, "user")
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
			str = append(str, results.Get(i).Field("uName").(string))
		}
	}

	return str
}
