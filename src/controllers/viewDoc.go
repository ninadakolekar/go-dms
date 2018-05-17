package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	models "github.com/ninadakolekar/aizant-dms/src/models"
	"github.com/rtt/Go-Solr"
)

//ViewDoc ... fetches document details
func ViewDoc(w http.ResponseWriter, r *http.Request) {

	fmt.Println("reached line no 10 in ViewDoc.go") //Debug
	tmpl := template.Must(template.ParseFiles("templates/viewDoc.html"))

	s, err := solr.Init("localhost", 8983, "docs2") //solr connection

	if err != nil {
		fmt.Println(err)
		return
	}
	quer := "id:1"
	q := solr.Query{
		Params: solr.URLParamMap{
			"q": []string{quer},
		},
		Rows: 1,
	}
	res, err := s.Select(&q)
	if err != nil {
		fmt.Println(err)
	}

	results := res.Results
	resss := models.Docs2{}
	resss.Init("1")
	fmt.Println(resss) //Debug
	if results.Len() != 0 {
		s1 := results.Get(0)
		resss.Template = s1.Field("template").(string)
		fmt.Println("template = ", resss.Template) //Debug
		temp, ok := s1.Field("paragraph").([]interface{})
		if ok {
			fmt.Println(temp) //Debug
			for i, v := range temp {
				fmt.Println("i = ", i) //Debug
				item, okk := v.(string)
				if okk {
					resss.Paragraph = append(resss.Paragraph, item)
				} else {
					break
				}
			}
		}
		fmt.Println(temp) //Debug
	}

	tmpl.Execute(w, resss)
}
