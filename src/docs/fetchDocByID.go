package docs

import (
	"errors"
	"fmt"

	"github.com/ninadakolekar/aizant-dms/src/constants"
	models "github.com/ninadakolekar/aizant-dms/src/models"
	solr "github.com/rtt/Go-Solr"
)

//FetchDocByID ... fetches
func FetchDocByID(uid string) (models.InactiveDoc, error) {

	var doc models.InactiveDoc
	s, err := solr.Init(constants.SolrHost, constants.SolrPort, constants.DocsCore)
	if err != nil {
		fmt.Println(err)
		return doc, err
	}
	q := solr.Query{
		Params: solr.URLParamMap{
			"q": []string{"id:" + uid},
		},
		Rows: 1,
	}
	res, err := s.Select(&q)

	if err != nil {
		fmt.Println(err)
		return doc, err
	}

	results := res.Results
	if results.Len() == 0 {
		err := errors.New("Document with id " + uid + "not found")
		return doc, err
	}
	s1 := results.Get(0)
	doc.DocNo = s1.Field("id").(string)
	//fmt.Println("it came to line 59") //Debug
	doc.Title = s1.Field("title").(string)
	//fmt.Println("it came to line 61") //Debug
	doc.DocType = s1.Field("docType").(string)
	//fmt.Println("it came to line 63") //Debug
	doc.DocProcess = s1.Field("docProcess").(string)
	doc.DocEffDate = s1.Field("effDate").(string)
	doc.DocExpDate = s1.Field("expDate").(string)
	// doc.DocProcess = s1.Field("docProcess").(string)
	doc.DocStatus = s1.Field("docStatus").(bool)
	//fmt.Println("it came to line 65") //Debug
	doc.Initiator = s1.Field("initiator").(string)
	//fmt.Println("it came to line 67") //Debug
	doc.Creator = s1.Field("creator").(string)
	//fmt.Println("it came to line 69") //Debug
	doc.DocDept = s1.Field("docDepartment").(string)
	//fmt.Println("it came to line 71") //Debug
	doc.FlowStatus = s1.Field("flowStatus").(float64)
	//fmt.Println("it came to line 73") //Debug
	doc.DocTemplate = s1.Field("docTemplateID").(float64)
	//fmt.Println("it came to line 75") //Debug
	// doc.DocTemplate = s1.Field("template").(string)
	doc.InitTS = s1.Field("initTime").(string)
	//fmt.Println("it came to line 77") //Debug

	createTime := s1.Field("createTime")
	if createTime == nil {
		doc.CreateTS = ""
	} else {
		doc.CreateTS = s1.Field("createTime").(string)
	}

	temp, ok := s1.Field("reviewer").([]interface{})
	if ok {
		for _, v := range temp {
			item, okk := v.(string)
			if okk {
				doc.Reviewer = append(doc.Reviewer, item)
			} else {
				break
			}
		}
	}
	temp, ok = s1.Field("approver").([]interface{})
	if ok {
		for _, v := range temp {
			item, okk := v.(string)
			if okk {
				doc.Approver = append(doc.Approver, item)
			} else {
				break
			}
		}
	}
	temp, ok = s1.Field("authorizer").([]interface{})
	if ok {
		for _, v := range temp {
			item, okk := v.(string)
			if okk {
				doc.Authorizer = append(doc.Authorizer, item)
			} else {
				break
			}
		}
	}
	return doc, nil
}
