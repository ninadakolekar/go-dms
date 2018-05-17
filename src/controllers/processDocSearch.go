package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	constant "github.com/ninadakolekar/aizant-dms/src/constants"
	solr "github.com/rtt/Go-Solr"
)

//ProcessDocSearch ... process doc search
func ProcessDocSearch(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		r.ParseForm()
		sCriteria := r.Form["criteria"][0]
		sKeyword := r.Form["searchKeyword"][0]
		fmt.Println("Form recieved : ", sCriteria, sKeyword, "\n ") //Debug
		links := []string{}
		boo := false
		if validateSearchForm(sCriteria) == true {
			query := makeSearchQuery([]string{sCriteria}, []string{sKeyword})
			fmt.Println("QUERY :: ", query)

			s, err := solr.Init(constant.SolrHost, constant.SolrPort, constant.DocsCore)
			if err != nil {
				fmt.Println("line number 25\n\n", err) //Debug

			}

			q := solr.Query{
				Params: solr.URLParamMap{
					"q": []string{query},
				},
				Rows: 100,
			}
			res, err := s.Select(&q)
			if err != nil {
				fmt.Println(err)
			}

			results := res.Results

			for i := 0; i < results.Len(); i++ {
				fmt.Println(results.Get(i), "\n\n") //Debug
				links = append(links, results.Get(i).Field("id").(string))
				boo = true
			}
			if results.Len() == 0 {
				fmt.Println("notfound\n\n") //Debug
			}

		} else {
			fmt.Println("not found\n\n") //Debug
		}
		tmpl := template.Must(template.ParseFiles("templates/searchDoc1.html"))
		tmpl.Execute(w, struct {
			Sarray []string
			Valid  bool
			Ivalid bool
		}{links, boo, !boo})
	} else {
		tmpl := template.Must(template.ParseFiles("templates/searchDoc1.html"))
		tmpl.Execute(w, struct{ temp bool }{true})
	}
}

func makeSearchQuery(sC []string, sK []string) string {
	validCriterion := []string{"docNumber", "docName", "docKeyword"}
	validQueryPrifex := []string{"id:", "title:", "body:"}

	querys := []string{}
	counter := 0
	for j, sc := range sC {
		fmt.Println(j, "\n\n its j") //Debug
		for i, v := range validCriterion {
			fmt.Println(i, "\n\n its i") //Debug
			if v == sc {
				fmt.Println("this came to if cond..") //Debug
				querys = append(querys, validQueryPrifex[i]+sK[j])
				fmt.Println(querys[counter])
				counter++
			}
		}
	}
	query := ""
	for i, q := range querys {
		if i == 0 {
			query = q
		} else {
			query += ("AND" + q)
		}
	}
	fmt.Println("\n\n line number 59") //Debug
	return query
}

func validateSearchForm(sC string) bool {
	validCriterion := []string{"docNumber", "docName", "docKeyword"}
	for _, v := range validCriterion {
		if v == sC {
			return true
		}
	}
	return false
}
