package docs

import (
	"errors"
	"fmt"

	"github.com/ninadakolekar/aizant-dms/src/constants"
	models "github.com/ninadakolekar/aizant-dms/src/models"
	solr "github.com/rtt/Go-Solr"
)

//FetchDocByID ... Fetches the document with given ID from the database
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

	if !isDocValid(s1) {
		err := errors.New("Invalid Document!\nMandatory fields found missing")
		return doc, err
	}

	doc.DocNo = s1.Field("id").(string)
	doc.Title = s1.Field("title").(string)
	doc.DocType = s1.Field("docType").(string)
	doc.DocProcess = s1.Field("docProcess").(string)
	doc.DocEffDate = s1.Field("effDate").(string)
	doc.DocExpDate = s1.Field("expDate").(string)
	doc.DocStatus = s1.Field("docStatus").(bool)
	doc.Initiator = s1.Field("initiator").(string)
	doc.Creator = s1.Field("creator").(string)
	doc.DocDept = s1.Field("docDepartment").(string)
	doc.FlowStatus = s1.Field("flowStatus").(float64)
	doc.DocTemplate = s1.Field("docTemplateID").(float64)
	doc.InitTS = s1.Field("initTime").(string)

	createTime := s1.Field("createTime")
	if createTime == nil {
		doc.CreateTS = ""
	} else {
		doc.CreateTS = s1.Field("createTime").(string)
	}

	reviewTime := s1.Field("reviewTime")
	if reviewTime == nil {
		doc.ReviewTS = ""
	} else {
		doc.ReviewTS = s1.Field("reviewTime").(string)
	}

	approveTime := s1.Field("approveTime")
	if approveTime == nil {
		doc.ApproveTS = ""
	} else {
		doc.ApproveTS = s1.Field("approveTime").(string)
	}

	authTime := s1.Field("authTime")
	if authTime == nil {
		doc.AuthTS = ""
	} else {
		doc.AuthTS = s1.Field("authTime").(string)
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

	if s1.Field("authorizer") != nil {
		temp, ok := s1.Field("authorizer").([]interface{})
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
	}

	return doc, nil
}

func isDocValid(s1 *solr.Document) bool {

	if s1.Field("id") != nil && s1.Field("title") != nil && s1.Field("docType") != nil && s1.Field("docProcess") != nil && s1.Field("effDate") != nil && s1.Field("expDate") != nil && s1.Field("docStatus") != nil && s1.Field("initiator") != nil && s1.Field("creator") != nil && s1.Field("docDepartment") != nil && s1.Field("flowStatus") != nil && s1.Field("docTemplateID") != nil && s1.Field("initTime") != nil && s1.Field("approver") != nil && s1.Field("reviewer") != nil {
		return true
	}

	return false
}
