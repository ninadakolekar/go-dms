package controllers

import (
	"fmt"
	"net/http"
	"time"

	auth "github.com/ninadakolekar/aizant-dms/src/auth"
	constant "github.com/ninadakolekar/aizant-dms/src/constants"
	solr "github.com/rtt/Go-Solr"
)

//FetchPendingDocuments ... fetch
func FetchPendingDocuments(w http.ResponseWriter, r *http.Request) {

	usr, err := auth.GetCurrentUser(r)
	if err != nil {
		fmt.Fprintf(w, "server error")
	}
	usr = "self"
	str1 := fetchInits(usr)
	str2 := fetchQA(usr)
	str3 := fetchCreator(usr)
	str4 := fetchReviews(usr)
	html := "<ul class='collapsible'><li><div class='collapsible-header'><i class='material-icons red'>place</i>Pending Intiations</div><div class='collapsible-body'><ul class='collection'>"
	html1 := "</ul></div></li><li><div class='collapsible-header'><i class='material-icons blue'>place</i>Pending QA</div><div class='collapsible-body'><ul  class='collection'>"
	html2 := "</ul></div></li><li><div class='collapsible-header'><i class='material-icons green'>place</i>Pending Create</div><div class='collapsible-body'><ul  class='collection'>"
	html3 := "</ul></div></li><li><div class='collapsible-header'><i class='material-icons indigo'>place</i>Pending Reviews</div><div class='collapsible-body'><ul  class='collection'>"

	html4 := "</ul></div></li><li><div class='collapsible-header'><i class='material-icons cyan'>place</i>Pending Approves</div><div class='collapsible-body'><ul  class='collection'>"
	html5 := "</ul></div></li><li><div class='collapsible-header'><i class='material-icons yellow'>place</i>Pending Authorises</div><div class='collapsible-body'><ul  class='collection'>"

	html6 := "</ul></div></li></ul><script>$(document).ready(function(){$('.collapsible').collapsible();});</script>"
	fmt.Fprintf(w, html+str1+html1+str2+html2+str3+html3+str4+html4+html5+html6)
}

func giveReviwerFlowStatus(results *solr.Document, usr string) bool {

	temp, ok := results.Field("reviewer").([]interface{})
	position := -1
	iter := 0
	reviewer := []string{}
	if ok {
		for _, v := range temp {
			item, okk := v.(string)
			if okk {
				iter++
				if item == usr {
					position = iter
				}
				reviewer = append(reviewer, item)
			} else {
				break
			}
		}
	}
	doctype := results.Field("docProcess").(string)
	flowstatus := results.Field("flowStatus").(float64)
	fmt.Println(flowstatus, position)
	if doctype == "Anyone" {
		return true
	} else if doctype == "Everyone" {
		return true
	}
	return true
}
func fetchReviews(usr string) string {
	query := "reviewer:" + usr
	s, err := solr.Init(constant.SolrHost, constant.SolrPort, constant.DocsCore)
	if err != nil {
		return "<h5>solr connection error</h5>"
	}
	q := solr.Query{
		Params: solr.URLParamMap{
			"q": []string{query},
		},
		Rows: 100,
	}

	res, err := s.Select(&q)
	if err != nil {
		return "<h5>solr connection error</h5>"
	}

	results := res.Results
	if results.Len() == 0 {
		return "<h5>No Pending Documents</h5>"
	}
	links := []lINK{}

	for i := 0; i < results.Len(); i++ {

		if giveReviwerFlowStatus(results.Get(i), usr) && results.Get(i).Field("docStatus").(bool) == false {
			links = append(links, convertTolINK(results.Get(i)))
		}
	}
	links = sortby(links, "expDate")

	r := ""
	now := time.Now()

	day := now.String()[0:10] + "T99:17:11.382Z"
	day3 := now.AddDate(0, 0, 3).String()[0:10] + "T99:17:11.382Z"
	day10 := now.AddDate(0, 0, 10).String()[0:10] + "T99:17:11.382Z"

	for _, e := range links {
		if e.ExpDate < day {
			r += ("<li class='collection-item avatar'><i class='material-icons circle #76ff03 black'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "Intiated &nbsp;<span class='fmtdate'>" + e.Idate + "</span></p><a href='" + "/doc/view/" + e.DocId + "' class = 'secondary-content'><i class='material-icons'>send</i></a></li>")
		} else if e.ExpDate < day3 {
			r += ("<li class='collection-item avatar'><i class='material-icons circle #76ff03 red'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "Intiated &nbsp;<span class='fmtdate'>" + e.Idate + "</span></p><a href='" + "/doc/view/" + e.DocId + "' class = 'secondary-content'><i class='material-icons'>send</i></a></li>")
		} else if e.ExpDate < day10 {
			r += ("<li class='collection-item avatar'><i class='material-icons circle #76ff03 pink'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "Intiated &nbsp;<span class='fmtdate'>" + e.Idate + "</span></p><a href='" + "/doc/view/" + e.DocId + "' class = 'secondary-content'><i class='material-icons'>send</i></a></li>")
		} else {
			r += ("<li class='collection-item avatar'><i class='material-icons circle #76ff03 grey'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "Intiated &nbsp;<span class='fmtdate'>" + e.Idate + "</span></p><a href='" + "/doc/view/" + e.DocId + "' class = 'secondary-content'><i class='material-icons'>send</i></a></li>")
		}
	}
	return r
}
func fetchCreator(usr string) string {
	query := "creator:" + usr
	s, err := solr.Init(constant.SolrHost, constant.SolrPort, constant.DocsCore)
	if err != nil {
		return "<h5>solr connection error</h5>"
	}
	q := solr.Query{
		Params: solr.URLParamMap{
			"q": []string{query},
		},
		Rows: 100,
	}

	res, err := s.Select(&q)
	if err != nil {
		return "<h5>solr connection error</h5>"
	}

	results := res.Results
	if results.Len() == 0 {
		return "<h5>No Pending Documents</h5>"
	}
	links := []lINK{}
	for i := 0; i < results.Len(); i++ {
		if results.Get(i).Field("flowStatus").(float64) == 2 && results.Get(i).Field("docStatus").(bool) == false {
			links = append(links, convertTolINK(results.Get(i)))
		}
	}
	links = sortby(links, "expDate")

	r := ""
	now := time.Now()

	day := now.String()[0:10] + "T99:17:11.382Z"
	day3 := now.AddDate(0, 0, 3).String()[0:10] + "T99:17:11.382Z"
	day10 := now.AddDate(0, 0, 10).String()[0:10] + "T99:17:11.382Z"

	for _, e := range links {
		if e.ExpDate < day {
			r += ("<li class='collection-item avatar'><i class='material-icons circle #76ff03 black'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "Intiated &nbsp;<span class='fmtdate'>" + e.Idate + "</span></p><a href='" + "/doc/view/" + e.DocId + "' class = 'secondary-content'><i class='material-icons'>send</i></a></li>")
		} else if e.ExpDate < day3 {
			r += ("<li class='collection-item avatar'><i class='material-icons circle #76ff03 red'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "Intiated &nbsp;<span class='fmtdate'>" + e.Idate + "</span></p><a href='" + "/doc/view/" + e.DocId + "' class = 'secondary-content'><i class='material-icons'>send</i></a></li>")
		} else if e.ExpDate < day10 {
			r += ("<li class='collection-item avatar'><i class='material-icons circle #76ff03 pink'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "Intiated &nbsp;<span class='fmtdate'>" + e.Idate + "</span></p><a href='" + "/doc/view/" + e.DocId + "' class = 'secondary-content'><i class='material-icons'>send</i></a></li>")
		} else {
			r += ("<li class='collection-item avatar'><i class='material-icons circle #76ff03 grey'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "Intiated &nbsp;<span class='fmtdate'>" + e.Idate + "</span></p><a href='" + "/doc/view/" + e.DocId + "' class = 'secondary-content'><i class='material-icons'>send</i></a></li>")
		}
	}
	return r
}

func fetchQA(usr string) string {
	query := "qa:" + usr
	s, err := solr.Init(constant.SolrHost, constant.SolrPort, constant.DocsCore)
	if err != nil {
		return "<h5>solr connection error</h5>"
	}
	q := solr.Query{
		Params: solr.URLParamMap{
			"q": []string{query},
		},
		Rows: 100,
	}

	res, err := s.Select(&q)
	if err != nil {
		return "<h5>solr connection error</h5>"
	}

	results := res.Results
	if results.Len() == 0 {
		return "<h5>No Pending Documents</h5>"
	}
	links := []lINK{}
	for i := 0; i < results.Len(); i++ {
		if results.Get(i).Field("flowStatus").(float64) == 1 && results.Get(i).Field("docStatus").(bool) == false {
			links = append(links, convertTolINK(results.Get(i)))
		}
	}
	links = sortby(links, "expDate")

	r := ""
	now := time.Now()

	day := now.String()[0:10] + "T99:17:11.382Z"
	day3 := now.AddDate(0, 0, 3).String()[0:10] + "T99:17:11.382Z"
	day10 := now.AddDate(0, 0, 10).String()[0:10] + "T99:17:11.382Z"

	for _, e := range links {
		if e.ExpDate < day {
			r += ("<li class='collection-item avatar'><i class='material-icons circle #76ff03 black'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "Intiated &nbsp;<span class='fmtdate'>" + e.Idate + "</span></p><a href='" + "/doc/view/" + e.DocId + "' class = 'secondary-content'><i class='material-icons'>send</i></a></li>")
		} else if e.ExpDate < day3 {
			r += ("<li class='collection-item avatar'><i class='material-icons circle #76ff03 red'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "Intiated &nbsp;<span class='fmtdate'>" + e.Idate + "</span></p><a href='" + "/doc/view/" + e.DocId + "' class = 'secondary-content'><i class='material-icons'>send</i></a></li>")
		} else if e.ExpDate < day10 {
			r += ("<li class='collection-item avatar'><i class='material-icons circle #76ff03 pink'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "Intiated &nbsp;<span class='fmtdate'>" + e.Idate + "</span></p><a href='" + "/doc/view/" + e.DocId + "' class = 'secondary-content'><i class='material-icons'>send</i></a></li>")
		} else {
			r += ("<li class='collection-item avatar'><i class='material-icons circle #76ff03 grey'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "Intiated &nbsp;<span class='fmtdate'>" + e.Idate + "</span></p><a href='" + "/doc/view/" + e.DocId + "' class = 'secondary-content'><i class='material-icons'>send</i></a></li>")
		}
	}
	return r
}
func fetchInits(usr string) string {
	query := "initiator:" + usr
	s, err := solr.Init(constant.SolrHost, constant.SolrPort, constant.DocsCore)
	if err != nil {
		return "<h5>solr connection error</h5>"
	}
	q := solr.Query{
		Params: solr.URLParamMap{
			"q": []string{query},
		},
		Rows: 100,
	}

	res, err := s.Select(&q)
	if err != nil {
		return "<h5>solr connection error</h5>"
	}

	results := res.Results
	if results.Len() == 0 {
		return "<h5>No Pending Documents</h5>"
	}
	links := []lINK{}
	for i := 0; i < results.Len(); i++ {
		if results.Get(i).Field("flowStatus").(float64) == 0 && results.Get(i).Field("docStatus").(bool) == false {
			links = append(links, convertTolINK(results.Get(i)))
		}
	}
	links = sortby(links, "expDate")

	r := ""
	now := time.Now()

	day := now.String()[0:10] + "T99:17:11.382Z"
	day3 := now.AddDate(0, 0, 3).String()[0:10] + "T99:17:11.382Z"
	day10 := now.AddDate(0, 0, 10).String()[0:10] + "T99:17:11.382Z"

	for _, e := range links {
		if e.ExpDate < day {
			r += ("<li class='collection-item avatar'><i class='material-icons circle #76ff03 black'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "Intiated &nbsp;<span class='fmtdate'>" + e.Idate + "</span></p><a href='" + "/doc/view/" + e.DocId + "' class = 'secondary-content'><i class='material-icons'>send</i></a></li>")
		} else if e.ExpDate < day3 {
			r += ("<li class='collection-item avatar'><i class='material-icons circle #76ff03 red'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "Intiated &nbsp;<span class='fmtdate'>" + e.Idate + "</span></p><a href='" + "/doc/view/" + e.DocId + "' class = 'secondary-content'><i class='material-icons'>send</i></a></li>")
		} else if e.ExpDate < day10 {
			r += ("<li class='collection-item avatar'><i class='material-icons circle #76ff03 pink'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "Intiated &nbsp;<span class='fmtdate'>" + e.Idate + "</span></p><a href='" + "/doc/view/" + e.DocId + "' class = 'secondary-content'><i class='material-icons'>send</i></a></li>")
		} else {
			r += ("<li class='collection-item avatar'><i class='material-icons circle #76ff03 grey'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "Intiated &nbsp;<span class='fmtdate'>" + e.Idate + "</span></p><a href='" + "/doc/view/" + e.DocId + "' class = 'secondary-content'><i class='material-icons'>send</i></a></li>")
		}
	}
	return r
}
