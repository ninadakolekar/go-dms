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
	fmt.Println(str1)
	html := "<ul class='collapsible'><li><div class='collapsible-header'><i class='material-icons'>filter_drama</i>PendingIntiations</div><div class='collapsible-body'><ul>"
	html1 := "</ul></div></li><li><div class='collapsible-header'><i class='material-icons'>place</i>PendingCreations</div><div class='collapsible-body'><ul>"
	html2 := "</ul></div></li><li><div class='collapsible-header'><i class='material-icons'>whatshot</i>PendingQA</div><div class='collapsible-body'><ul>"
	html3 := "</ul></div></li></ul><script>$(document).ready(function(){$('.collapsible').collapsible();});</script>"
	fmt.Fprintf(w, html+str1+html1+html2+html3)
}

func fetchInits(usr string) string {
	query := "initiator:" + usr //only InActive Documents.
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
		if results.Get(i).Field("flowStatus") == 0 {
			links = append(links, convertTolINK(results.Get(i)))
		}
	}
	fmt.Println("links at final", links)
	links = sortby(links, "expDate")

	r := ""
	now := time.Now()

	day := now.String()[0:10] + "T90.00.00Z" //CheckthisFormat
	day3 := now.AddDate(0, 0, 3).String()[0:10] + "T90.00.00Z"
	day10 := now.AddDate(0, 0, 10).String()[0:10] + "T90.00.00Z"

	for _, e := range links {
		if e.ExpDate < day {
			r += ("<li class='collection-item avatar'><i class='material-icons circle #76ff03 black'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "Intiated &nbsp;<span class='fmtdate'>" + e.Idate + "</span></p><a href='" + "/doc/view/" + e.DocId + "' class = 'secondary-content'><br><i class='material-icons'>send</i></a></li>")
		} else if e.ExpDate < day3 {
			r += ("<li class='collection-item avatar'><i class='material-icons circle #76ff03 red'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "Intiated &nbsp;<span class='fmtdate'>" + e.Idate + "</span></p><a href='" + "/doc/view/" + e.DocId + "' class = 'secondary-content'><br><i class='material-icons'>send</i></a></li>")
		} else if e.ExpDate < day10 {
			r += ("<li class='collection-item avatar'><i class='material-icons circle #76ff03 pink'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "Intiated &nbsp;<span class='fmtdate'>" + e.Idate + "</span></p><a href='" + "/doc/view/" + e.DocId + "' class = 'secondary-content'><br><i class='material-icons'>send</i></a></li>")
		} else {
			r += ("<li class='collection-item avatar'><i class='material-icons circle #76ff03 white'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "Intiated &nbsp;<span class='fmtdate'>" + e.Idate + "</span></p><a href='" + "/doc/view/" + e.DocId + "' class = 'secondary-content'><br><i class='material-icons'>send</i></a></li>")
		}
	}
	return r
}
