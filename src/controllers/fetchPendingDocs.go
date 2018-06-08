package controllers

import (
	"fmt"
	"html"
	"net/http"
	"strconv"
	"time"

	auth "github.com/ninadakolekar/go-dms/src/auth"
	constant "github.com/ninadakolekar/go-dms/src/constants"
	"github.com/ninadakolekar/go-dms/src/docs"
	solr "github.com/rtt/Go-Solr"
)

//FetchPendingDocuments ... fetch
func FetchPendingDocuments(w http.ResponseWriter, r *http.Request) {

	_, err := auth.GetCurrentUser(r)
	if err != nil {
		fmt.Fprintf(w, "server error")
	}
	//fmt.Println(makeConstraints(r))
	str1, i1 := fetchInits(r)
	str2, i2 := fetchQA(r)
	str3, i3 := fetchCreator(r)
	str4, i4 := fetchReviews(r)
	str5, i5 := fetchApproves(r)
	str6, i6 := fetchAuthorises(r)

	html := "<ul class='collapsible'><li><div class='collapsible-header'><i class='material-icons red'>layers</i>Pending Intiations <span class = 'docsCount'>" + strconv.Itoa(i1) + "</span></div><div class='collapsible-body'><ul class='collection'>"
	html1 := "</ul></div></li><li><div class='collapsible-header'><i class='material-icons blue'>layers</i>Pending QA <span class = 'docsCount'>" + strconv.Itoa(i2) + "</span></div><div class='collapsible-body'><ul  class='collection'>"
	html2 := "</ul></div></li><li><div class='collapsible-header'><i class='material-icons green'>layers</i>Pending Creations <span class = 'docsCount'>" + strconv.Itoa(i3) + "</span></div><div class='collapsible-body'><ul  class='collection'>"
	html3 := "</ul></div></li><li><div class='collapsible-header'><i class='material-icons indigo'>layers</i>Pending Reviews <span class = 'docsCount'>" + strconv.Itoa(i4) + "</span></div><div class='collapsible-body'><ul  class='collection'>"

	html4 := "</ul></div></li><li><div class='collapsible-header'><i class='material-icons cyan'>layers</i>Pending Approvals <span class = 'docsCount'>" + strconv.Itoa(i5) + "</span></div><div class='collapsible-body'><ul  class='collection'>"
	html5 := "</ul></div></li><li><div class='collapsible-header'><i class='material-icons yellow'>layers</i>Pending Authorizations <span class = 'docsCount'>" + strconv.Itoa(i6) + "</span></div><div class='collapsible-body'><ul  class='collection'>"

	html6 := "</ul></div></li></ul><script>$(document).ready(function(){$('.collapsible').collapsible();});</script>"
	fmt.Fprintf(w, html+str1+html1+str2+html2+str3+html3+str4+html4+str5+html5+str6+html6)
}

func makeConstraints(r *http.Request) string {
	sop := html.EscapeString(r.FormValue("SOP"))
	hr := html.EscapeString(r.FormValue("HR"))
	stp := html.EscapeString(r.FormValue("STP"))
	everyone := html.EscapeString(r.FormValue("Everyone"))
	onebyone := html.EscapeString(r.FormValue("OneByOne"))
	anyone := html.EscapeString(r.FormValue("Anyone"))
	d1 := html.EscapeString(r.FormValue("option1"))
	d2 := html.EscapeString(r.FormValue("option2"))
	d3 := html.EscapeString(r.FormValue("option3"))
	query := ""
	open := false
	if sop == "on" {
		query += "( docType:SOP "
		open = true
	}
	if hr == "on" {
		if query == "" {
			query += "( docType:HR "
			open = true
		} else {
			query += " docType:HR "
		}
	}
	if stp == "on" {
		if query == "" {
			open = true
			query += "( docType:STP "
		} else {
			query += " docType:STP "
		}
	}
	if open {
		query += " )"
		open = false
	}
	if everyone == "on" {
		if query != "" {
			query += " AND "
		}
		open = true
		query += "( docProcess:Everyone "
	}
	if onebyone == "on" {
		if query == "" {
			open = true
			query += "( docProcess:OneByOne "
		} else {
			if open {
				query += " docProcess:OneByOne "
			} else {
				open = true
				query += "( docProcess:OneByOne "
			}
		}
	}
	if anyone == "on" {
		if query == "" {
			open = true
			query += "(  docProcess:Anyone "
		} else {
			if open {
				query += "  docProcess:Anyone "
			} else {
				open = true
				query += "(  docProcess:Anyone "
			}
		}
	}
	if open {
		query += " )"
		open = false
	}
	if d1 == "on" {
		if query != "" {
			query += " AND "
		}
		open = true
		query += "( docDepartment:D1"
	}
	if d2 == "on" {
		if query == "" {
			open = true
			query += "( docDepartment:D2 "
		} else {
			if open {
				query += " docDepartment:D2 "
			} else {
				open = true
				query += "( docDepartment:D2 "
			}
		}
	}
	if d3 == "on" {
		if query == "" {
			open = true
			query += "(  docDepartment:D3 "
		} else {
			if open {
				query += "  docDepartment:D3 "
			} else {
				open = true
				query += "(  docDepartment:D3 "
			}
		}
	}
	if open {
		query += " )"
		open = false
	}
	return query
}
func giveFlowStatus(results *solr.Document, usr string, value int) bool {
	doc := docs.ConvertSolrDoc2InactiveModel(results)
	if value == 0 {
		return docs.CheckCurrentReviewer(doc, usr)
	} else if value == 1 {
		return docs.CheckCurrentApprover(doc, usr)
	} else if value == 2 {
		return docs.CheckCurrentAuthorizer(doc, usr)
	}
	return false
}
func fetchApproves(rr *http.Request) (string, int) {
	usr, _ := auth.GetCurrentUser(rr)

	mk := makeConstraints(rr)
	query := "approver:" + usr
	if mk != "" {
		query += (" AND " + mk)
	}
	//fmt.Println("QUERY:#", query, "#")
	s, err := solr.Init(constant.SolrHost, constant.SolrPort, constant.DocsCore)
	if err != nil {
		return "<h5>solr connection error</h5>", 0
	}
	q := solr.Query{
		Params: solr.URLParamMap{
			"q": []string{query},
		},
		Rows: 100,
	}

	res, err := s.Select(&q)
	if err != nil {
		return "<h5>solr connection error</h5>", 0
	}

	results := res.Results
	if results.Len() == 0 {
		return "<h5>No Pending Documents</h5>", 0
	}
	links := []lINK{}

	for i := 0; i < results.Len(); i++ {

		if giveFlowStatus(results.Get(i), usr, 1) && results.Get(i).Field("docStatus").(bool) == false {
			links = append(links, convertTolINK(results.Get(i)))
		}
	}
	links = sortby(links, "expDate")
	if len(links) == 0 {
		return "<h5>No Pending Documents</h5>", 0
	}
	r := ""
	lenn := len(links)
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
	r += "<script>$('.fmtdate').each(function(){var date = $(this).html();var formattedDate = date.split('T')[0];var fDate = formattedDate.split('-');$(this).html(fDate[2]+'-'+fDate[1]+'-'+fDate[0]);})</script>"
	return r, lenn
}
func fetchAuthorises(rr *http.Request) (string, int) {
	usr, _ := auth.GetCurrentUser(rr)

	mk := makeConstraints(rr)
	query := "authorizer:" + usr
	if mk != "" {
		query += (" AND " + mk)
	}
	//fmt.Println("QUERY:#", query, "#")
	s, err := solr.Init(constant.SolrHost, constant.SolrPort, constant.DocsCore)
	if err != nil {
		return "<h5>solr connection error</h5>", 0
	}
	q := solr.Query{
		Params: solr.URLParamMap{
			"q": []string{query},
		},
		Rows: 100,
	}

	res, err := s.Select(&q)
	if err != nil {
		return "<h5>solr connection error</h5>", 0
	}

	results := res.Results
	if results.Len() == 0 {
		return "<h5>No Pending Documents</h5>", 0
	}
	links := []lINK{}

	for i := 0; i < results.Len(); i++ {

		if giveFlowStatus(results.Get(i), usr, 2) && results.Get(i).Field("docStatus").(bool) == false {
			links = append(links, convertTolINK(results.Get(i)))
		}
	}
	links = sortby(links, "expDate")
	if len(links) == 0 {
		return "<h5>No Pending Documents</h5>", 0
	}
	r := ""
	lenn := len(links)
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
	r += "<script>$('.fmtdate').each(function(){var date = $(this).html();var formattedDate = date.split('T')[0];var fDate = formattedDate.split('-');$(this).html(fDate[2]+'-'+fDate[1]+'-'+fDate[0]);})</script>"
	return r, lenn
}

func fetchReviews(rr *http.Request) (string, int) {
	usr, _ := auth.GetCurrentUser(rr)

	mk := makeConstraints(rr)
	query := "reviewer:" + usr
	if mk != "" {
		query += (" AND " + mk)
	}
	//fmt.Println("QUERY:#", query, "#")
	s, err := solr.Init(constant.SolrHost, constant.SolrPort, constant.DocsCore)
	if err != nil {
		return "<h5>solr connection error</h5>", 0
	}
	q := solr.Query{
		Params: solr.URLParamMap{
			"q": []string{query},
		},
		Rows: 100,
	}

	res, err := s.Select(&q)
	if err != nil {
		return "<h5>solr connection error</h5>", 0
	}

	results := res.Results
	if results.Len() == 0 {
		return "<h5>No Pending Documents</h5>", 0
	}
	links := []lINK{}

	for i := 0; i < results.Len(); i++ {

		if giveFlowStatus(results.Get(i), usr, 0) && results.Get(i).Field("docStatus").(bool) == false {
			links = append(links, convertTolINK(results.Get(i)))
		}
	}
	links = sortby(links, "expDate")
	if len(links) == 0 {
		return "<h5>No Pending Documents</h5>", 0
	}
	r := ""
	lenn := len(links)
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
	r += "<script>$('.fmtdate').each(function(){var date = $(this).html();var formattedDate = date.split('T')[0];var fDate = formattedDate.split('-');$(this).html(fDate[2]+'-'+fDate[1]+'-'+fDate[0]);})</script>"
	return r, lenn
}

func fetchCreator(rr *http.Request) (string, int) {
	usr, _ := auth.GetCurrentUser(rr)

	mk := makeConstraints(rr)
	query := "creator:" + usr
	if mk != "" {
		query += (" AND " + mk)
	}
	//fmt.Println("QUERY:#", query, "#")
	s, err := solr.Init(constant.SolrHost, constant.SolrPort, constant.DocsCore)
	if err != nil {
		return "<h5>solr connection error</h5>", 0
	}
	q := solr.Query{
		Params: solr.URLParamMap{
			"q": []string{query},
		},
		Rows: 100,
	}

	res, err := s.Select(&q)
	if err != nil {
		return "<h5>solr connection error</h5>", 0
	}

	results := res.Results
	if results.Len() == 0 {
		return "<h5>No Pending Documents</h5>", 0
	}
	links := []lINK{}
	for i := 0; i < results.Len(); i++ {
		if results.Get(i).Field("flowStatus").(float64) == 2 && results.Get(i).Field("docStatus").(bool) == false {
			links = append(links, convertTolINK(results.Get(i)))
		}
	}
	links = sortby(links, "expDate")
	if len(links) == 0 {
		return "<h5>No Pending Documents</h5>", 0
	}
	r := ""
	lenn := len(links)
	now := time.Now()

	day := now.String()[0:10] + "T99:17:11.382Z"
	day3 := now.AddDate(0, 0, 3).String()[0:10] + "T99:17:11.382Z"
	day10 := now.AddDate(0, 0, 10).String()[0:10] + "T99:17:11.382Z"

	for _, e := range links {
		if e.ExpDate < day {
			r += ("<li class='collection-item avatar'><i class='material-icons circle #76ff03 black'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "Intiated &nbsp;<span class='fmtdate'>" + e.Idate + "</span></p><a href='" + "/doc/create/" + e.DocId + "' class = 'secondary-content'><i class='material-icons'>send</i></a></li>")
		} else if e.ExpDate < day3 {
			r += ("<li class='collection-item avatar'><i class='material-icons circle #76ff03 red'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "Intiated &nbsp;<span class='fmtdate'>" + e.Idate + "</span></p><a href='" + "/doc/create/" + e.DocId + "' class = 'secondary-content'><i class='material-icons'>send</i></a></li>")
		} else if e.ExpDate < day10 {
			r += ("<li class='collection-item avatar'><i class='material-icons circle #76ff03 pink'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "Intiated &nbsp;<span class='fmtdate'>" + e.Idate + "</span></p><a href='" + "/doc/create/" + e.DocId + "' class = 'secondary-content'><i class='material-icons'>send</i></a></li>")
		} else {
			r += ("<li class='collection-item avatar'><i class='material-icons circle #76ff03 grey'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "Intiated &nbsp;<span class='fmtdate'>" + e.Idate + "</span></p><a href='" + "/doc/create/" + e.DocId + "' class = 'secondary-content'><i class='material-icons'>send</i></a></li>")
		}
	}
	r += "<script>$('.fmtdate').each(function(){var date = $(this).html();var formattedDate = date.split('T')[0];var fDate = formattedDate.split('-');$(this).html(fDate[2]+'-'+fDate[1]+'-'+fDate[0]);})</script>"
	return r, lenn
}

func fetchQA(rr *http.Request) (string, int) {
	usr, _ := auth.GetCurrentUser(rr)

	mk := makeConstraints(rr)
	query := "qa:" + usr
	if mk != "" {
		query += (" AND " + mk)
	}
	//fmt.Println("QUERY:#", query, "#")
	s, err := solr.Init(constant.SolrHost, constant.SolrPort, constant.DocsCore)
	if err != nil {
		return "<h5>solr connection error</h5>", 0
	}
	q := solr.Query{
		Params: solr.URLParamMap{
			"q": []string{query},
		},
		Rows: 100,
	}

	res, err := s.Select(&q)
	if err != nil {
		return "<h5>solr connection error</h5>", 0
	}

	results := res.Results
	if results.Len() == 0 {
		return "<h5>No Pending Documents</h5>", 0
	}
	links := []lINK{}
	for i := 0; i < results.Len(); i++ {
		if results.Get(i).Field("flowStatus").(float64) == 1 && results.Get(i).Field("docStatus").(bool) == false {
			links = append(links, convertTolINK(results.Get(i)))
		}
	}
	links = sortby(links, "expDate")
	lenn := len(links)

	if len(links) == 0 {
		return "<h5>No Pending Documents</h5>", 0
	}
	r := ""
	now := time.Now()

	day := now.String()[0:10] + "T99:17:11.382Z"
	day3 := now.AddDate(0, 0, 3).String()[0:10] + "T99:17:11.382Z"
	day10 := now.AddDate(0, 0, 10).String()[0:10] + "T99:17:11.382Z"

	for _, e := range links {
		if e.ExpDate < day {
			r += ("<li class='collection-item avatar'><i class='material-icons circle #76ff03 black'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "Intiated &nbsp;<span class='fmtdate'>" + e.Idate + "</span></p><a href='" + "/doc/viewDetails/" + e.DocId + "' class = 'secondary-content'><i class='material-icons'>send</i></a></li>")
		} else if e.ExpDate < day3 {
			r += ("<li class='collection-item avatar'><i class='material-icons circle #76ff03 red'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "Intiated &nbsp;<span class='fmtdate'>" + e.Idate + "</span></p><a href='" + "/doc/viewDetails/" + e.DocId + "' class = 'secondary-content'><i class='material-icons'>send</i></a></li>")
		} else if e.ExpDate < day10 {
			r += ("<li class='collection-item avatar'><i class='material-icons circle #76ff03 pink'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "Intiated &nbsp;<span class='fmtdate'>" + e.Idate + "</span></p><a href='" + "/doc/viewDetails/" + e.DocId + "' class = 'secondary-content'><i class='material-icons'>send</i></a></li>")
		} else {
			r += ("<li class='collection-item avatar'><i class='material-icons circle #76ff03 grey'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "Intiated &nbsp;<span class='fmtdate'>" + e.Idate + "</span></p><a href='" + "/doc/viewDetails/" + e.DocId + "' class = 'secondary-content'><i class='material-icons'>send</i></a></li>")
		}
	}
	r += "<script>$('.fmtdate').each(function(){var date = $(this).html();var formattedDate = date.split('T')[0];var fDate = formattedDate.split('-');$(this).html(fDate[2]+'-'+fDate[1]+'-'+fDate[0]);})</script>"
	return r, lenn
}
func fetchInits(rr *http.Request) (string, int) {
	usr, _ := auth.GetCurrentUser(rr)

	mk := makeConstraints(rr)
	query := "initiator:" + usr
	if mk != "" {
		query += (" AND " + mk)
	}
	//fmt.Println("QUERY:#", query, "#")
	s, err := solr.Init(constant.SolrHost, constant.SolrPort, constant.DocsCore)
	if err != nil {
		return "<h5>solr connection error</h5>", 0
	}
	q := solr.Query{
		Params: solr.URLParamMap{
			"q": []string{query},
		},
		Rows: 100,
	}

	res, err := s.Select(&q)
	if err != nil {
		return "<h5>solr connection error</h5>", 0
	}

	results := res.Results
	if results.Len() == 0 {
		return "<h5>No Pending Documents</h5>", 0
	}
	links := []lINK{}
	for i := 0; i < results.Len(); i++ {
		if results.Get(i).Field("flowStatus").(float64) == 0 && results.Get(i).Field("docStatus").(bool) == false {
			links = append(links, convertTolINK(results.Get(i)))
		}
	}

	lenn := len(links)
	links = sortby(links, "expDate")
	if len(links) == 0 {
		return "<h5>No Pending Documents</h5>", lenn
	}
	r := ""
	now := time.Now()

	day := now.String()[0:10] + "T99:17:11.382Z"
	day3 := now.AddDate(0, 0, 3).String()[0:10] + "T99:17:11.382Z"
	day10 := now.AddDate(0, 0, 10).String()[0:10] + "T99:17:11.382Z"

	for _, e := range links {
		if e.ExpDate < day {
			r += ("<li class='collection-item avatar'><i class='material-icons circle #76ff03 black'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "Intiated &nbsp;<span class='fmtdate'>" + e.Idate + "</span></p><a href='" + "/doc/add/" + e.DocId + "' class = 'secondary-content'><i class='material-icons'>send</i></a></li>")
		} else if e.ExpDate < day3 {
			r += ("<li class='collection-item avatar'><i class='material-icons circle #76ff03 red'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "Intiated &nbsp;<span class='fmtdate'>" + e.Idate + "</span></p><a href='" + "/doc/add/" + e.DocId + "' class = 'secondary-content'><i class='material-icons'>send</i></a></li>")
		} else if e.ExpDate < day10 {
			r += ("<li class='collection-item avatar'><i class='material-icons circle #76ff03 pink'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "Intiated &nbsp;<span class='fmtdate'>" + e.Idate + "</span></p><a href='" + "/doc/add/" + e.DocId + "' class = 'secondary-content'><i class='material-icons'>send</i></a></li>")
		} else {
			r += ("<li class='collection-item avatar'><i class='material-icons circle #76ff03 grey'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "Intiated &nbsp;<span class='fmtdate'>" + e.Idate + "</span></p><a href='" + "/doc/add/" + e.DocId + "' class = 'secondary-content'><i class='material-icons'>send</i></a></li>")
		}
	}
	r += "<script>$('.fmtdate').each(function(){var date = $(this).html();var formattedDate = date.split('T')[0];var fDate = formattedDate.split('-');$(this).html(fDate[2]+'-'+fDate[1]+'-'+fDate[0]);})</script>"
	return r, lenn
}
