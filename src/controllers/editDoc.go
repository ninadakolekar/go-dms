package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	models "github.com/ninadakolekar/aizant-dms/src/models"
	"github.com/rtt/Go-Solr"
)

//EditDoc ... edit document details
func EditDoc(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/editDoc.html"))

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

	temp := convertToDocs2(s.Select(&q))

	tmpl.Execute(w, temp)
}

//function that converts doc2SelectResponse, error schema to model doc2
func convertToDocs2(si *solr.SelectResponse, err error) models.Docs2 {

	res, err := si, err

	var doc2 models.Docs2
	doc2.Init("%$%")

	if err != nil {
		fmt.Println(err)
		return doc2

	}

	results := res.Results

	if results.Len() != 0 {
		//fmt.Println("it came to line 55") //Debug
		s1 := results.Get(0)
		//fmt.Println("it came to line 57") //Debug
		doc2.DocNo = s1.Field("id").(string)
		//fmt.Println("it came to line 59") //Debug
		doc2.Title = s1.Field("title").(string)
		//fmt.Println("it came to line 61") //Debug
		doc2.DocType = s1.Field("docType").(string)
		//fmt.Println("it came to line 63") //Debug
		doc2.DocStatus = s1.Field("docStatus").(bool)
		//fmt.Println("it came to line 65") //Debug
		doc2.Initiator = s1.Field("initiator").(string)
		//fmt.Println("it came to line 67") //Debug
		doc2.Creator = s1.Field("creator").(string)
		//fmt.Println("it came to line 69") //Debug
		doc2.DocDept = s1.Field("docDepartment").(string)
		//fmt.Println("it came to line 71") //Debug
		doc2.FlowStatus = s1.Field("flowStatus").(float64)
		//fmt.Println("it came to line 73") //Debug
		doc2.DocTemplate = s1.Field("docTemplateID").(float64)
		//fmt.Println("it came to line 75") //Debug
		doc2.Template = s1.Field("template").(string)
		//fmt.Println("it came to line 77") //Debug
		temp, ok := s1.Field("reviewer").([]interface{})
		if ok {
			for _, v := range temp {
				item, okk := v.(string)
				if okk {
					doc2.Reviewer = append(doc2.Reviewer, item)
				} else {
					break
				}
			}
		}
		//fmt.Println("\n\n line no 76") //Debug
		temp, ok = s1.Field("approver").([]interface{})
		if ok {
			for _, v := range temp {
				item, okk := v.(string)
				if okk {
					doc2.Approver = append(doc2.Approver, item)
				} else {
					break
				}
			}
		}
		//fmt.Println("\n\n line no 88") //Debug
		temp, ok = s1.Field("authorizer").([]interface{})
		if ok {
			for _, v := range temp {
				item, okk := v.(string)
				if okk {
					doc2.Authorizer = append(doc2.Authorizer, item)
				} else {
					break
				}
			}
		}
		//fmt.Println("this ") //debug
		temp, ok = s1.Field("paragraph").([]interface{})
		if ok {
			for _, v := range temp {
				item, okk := v.(string)
				if okk {
					doc2.Paragraph = append(doc2.Paragraph, item)
				} else {
					break
				}
			}
		}
		temp, ok = s1.Field("images").([]interface{})
		if ok {
			for _, v := range temp {
				item, okk := v.(string)
				if okk {
					doc2.Images = append(doc2.Images, item)
				} else {
					break
				}
			}
		}
		temp, ok = s1.Field("tables").([]interface{})
		if ok {
			for _, v := range temp {
				item, okk := v.(string)
				if okk {
					doc2.Tables = append(doc2.Tables, item)
				} else {
					break
				}
			}
		}

	}
	//fmt.Println("\n\n line no 133")
	return doc2
}
