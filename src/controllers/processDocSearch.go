package controllers

import (
	"fmt"
	"html"
	"net/http"
	"regexp"
	"strings"

	constant "github.com/ninadakolekar/aizant-dms/src/constants"
	solr "github.com/rtt/Go-Solr"
)

//ProcessDocSearch ... process new search
func ProcessDocSearch(w http.ResponseWriter, r *http.Request) {

	links := []lINK{}
	sortOrder := "typeSort" //"alexical" //"default"
	query := buildQuery(r)
	fmt.Println("query:", query)
	if query == "" {
		fmt.Fprintf(w, "<div class='card-panel'><span class='blue-text text-darken-2'><h3>Invalid query.</h3></span></div>")
	} else if query == "empty" {
		fmt.Fprintf(w, "<div class='card-panel'><span class='blue-text text-darken-2'><h3></h3></span></div>")
	} else {
		s, err := solr.Init(constant.SolrHost, constant.SolrPort, constant.DocsCore)
		if err != nil {
			fmt.Println(err) //core not connected
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
		if results.Len() == 0 {
			fmt.Fprintf(w, "<div class='card-panel'><span class='blue-text text-darken-2'><h5>No results found.</h5></span></div>")
		} else {
			for i := 0; i < results.Len(); i++ {
				links = append(links, convertTolINK(results.Get(i)))
			}

			links = sortby(links, sortOrder)
			if sortOrder == "typeSort" {

				s1 := "<ul class='collapsible'><li><div class='collapsible-header'><i class='material-icons circle #76ff03 red'>insert_drive_file</i>HR</div><div class='collapsible-body'><ul>"
				v1 := ""
				v2 := ""
				v3 := ""
				for _, e := range links {
					if e.DocType == "HR" {
						v1 += "<li class='collection-item avatar'><i class='material-icons circle #76ff03 red'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "intiated:" + e.Idate + "</p><a href='" + "/doc/view/" + e.DocId + "' class = 'secondary-content'><br><i class='material-icons'>send</i></a></li>"
					} else if e.DocType == "SOP" {
						v2 += "<li class='collection-item avatar'><i class='material-icons circle #76ff03 blue'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "intiated:" + e.Idate + "</p><a href='" + "/doc/view/" + e.DocId + "' class = 'secondary-content'><br><i class='material-icons'>send</i></a></li>"
					} else if e.DocType == "STP" {
						v3 += "<li class='collection-item avatar'><i class='material-icons circle #76ff03 green'>insert_drive_file</i><span class='title'>" + e.DocName + "</span><p>" + "intiated:" + e.Idate + "</p><a href='" + "/doc/view/" + e.DocId + "' class = 'secondary-content'><br><i class='material-icons'>send</i></a></li>"
					}
				}

				s2 := "</ul></div></li><li><div class='collapsible-header'><i class='material-icons circle #76ff03 blue'>insert_drive_file</i>SOP</div><div class='collapsible-body'><ul>"
				s3 := "</ul></div></li><li><div class='collapsible-header'><i class='material-icons circle #76ff03 green'>insert_drive_file</i>STP</div><div class='collapsible-body'><ul>"
				s4 := "</ul></div></li></ul><script>$(document).ready(function(){$('.collapsible').collapsible();});</script>"
				fmt.Fprintf(w, s1+v1+s2+v2+s3+v3+s4)
			} else {
				s1 := "<li class='collection-item avatar'><i class='material-icons circle #76ff03 "
				color := "green"
				s11 := "'>insert_drive_file</i><span class='title'>"
				s2 := "</span><p>"
				s3 := "</p><a href='"
				s4 := "' class = 'secondary-content'><br><i class='material-icons'>send</i></a></li>"

				ret := ""
				for _, e := range links {
					if e.DocType == "SOP" {
						color = "blue"
					} else if e.DocType == "HR" {
						color = "red"
					} else {
						color = "green"
					}
					ret += (s1 + color + s11 + e.DocName + s2 + "intiated:" + e.Idate + s3 + "/doc/view/" + e.DocId + s4)
				}

				fmt.Fprintf(w, ret)
			}

		}
	}

}

type formData struct {
	sop      string
	hr       string
	stp      string
	everyone string
	onebyone string
	anyone   string
	d1       string
	d2       string
	d3       string
	initFrom string
	initTo   string
	effFrom  string
	effTo    string
	expFrom  string
	expTo    string
	select1  string
	val1     string
	select2  string
	val2     string
	select3  string
	val3     string
}

func buildQuery(r *http.Request) string { //never returns an empty string

	information := formData{
		sop:      html.EscapeString(r.FormValue("SOP")),
		hr:       html.EscapeString(r.FormValue("HR")),
		stp:      html.EscapeString(r.FormValue("STP")),
		everyone: html.EscapeString(r.FormValue("Everyone")),
		onebyone: html.EscapeString(r.FormValue("OneByOne")),
		anyone:   html.EscapeString(r.FormValue("Anyone")),
		d1:       html.EscapeString(r.FormValue("option1")),
		d2:       html.EscapeString(r.FormValue("option2")),
		d3:       html.EscapeString(r.FormValue("option3")),
		initFrom: html.EscapeString(r.FormValue("initFrom")),
		initTo:   html.EscapeString(r.FormValue("initTo")),
		effFrom:  html.EscapeString(r.FormValue("effFrom")),
		effTo:    html.EscapeString(r.FormValue("effTo")),
		expFrom:  html.EscapeString(r.FormValue("expFrom")),
		expTo:    html.EscapeString(r.FormValue("expTo")),
		select1:  html.EscapeString(r.FormValue("select1")),
		val1:     html.EscapeString(r.FormValue("val1")),
		select2:  html.EscapeString(r.FormValue("select2")),
		val2:     html.EscapeString(r.FormValue("val2")),
		select3:  html.EscapeString(r.FormValue("select3")),
		val3:     html.EscapeString(r.FormValue("val3")),
	}

	if !isEmpty(information) {
		return "empty"
	}
	if !validate(information) {
		return ""
	}

	query := makedateQuery(information)
	typeQuery := makeTypeQuery(information)
	processQuery := makeProcessQuery(information)
	departmentQuery := makeDepartmentQuery(information)
	criteraQuery := makeCritriaQuery(information)
	if typeQuery != "" {
		query += (" AND (" + typeQuery + ")")
	}
	if processQuery != "" {
		query += (" AND (" + processQuery + ")")
	}
	if departmentQuery != "" {
		query += (" AND (" + departmentQuery + ")")
	}
	if criteraQuery != "" {
		query += (" AND (" + criteraQuery + ")")
	}
	return query
}

func makeCritriaQuery(x formData) string {
	query := ""
	if x.select1 != "" {
		q1 := ""
		if x.select1 == "docNumber" {
			q1 += ("id:(" + x.val1 + "* )")
		} else if x.select1 == "docTitle" {
			q1 += ("title:(" + strings.Replace(x.val1, " ", "*", -1) + "* )")
		} else if x.select1 == "docContent" {
			q1 += ("body:( *" + strings.Replace(x.val1, " ", "*", -1) + "* )")
		} else if x.select1 == "docUser" {
			q1 += (" creator:(" + strings.Replace(x.val1, " ", "*", -1) + ") approver:(" + strings.Replace(x.val1, " ", "*", -1) + ") authorizer:(" + strings.Replace(x.val1, " ", "*", -1) + ") initiator:(" + strings.Replace(x.val1, " ", "*", -1) + ") reviewer:(" + strings.Replace(x.val1, " ", "*", -1) + ")")
		}
		query += ("(" + q1 + ")")
	}
	if x.select2 != "" {
		q1 := ""
		if x.select2 == "docNumber" {
			q1 += ("id:(" + x.val2 + "* )")
		} else if x.select2 == "docTitle" {
			q1 += ("title:(" + strings.Replace(x.val2, " ", "*", -1) + "* )")
		} else if x.select2 == "docContent" {
			q1 += ("body:( *" + strings.Replace(x.val2, " ", "*", -1) + "* )")
		} else if x.select2 == "docUser" {
			q1 += (" creator:(" + strings.Replace(x.val2, " ", "*", -1) + ") approver:(" + strings.Replace(x.val2, " ", "*", -1) + ") authorizer:(" + strings.Replace(x.val2, " ", "*", -1) + ") initiator:(" + strings.Replace(x.val2, " ", "*", -1) + ") reviewer:(" + strings.Replace(x.val2, " ", "*", -1) + ")")
		}
		if query == "" {
			query += ("(" + q1 + ")")
		} else {
			query += (" AND (" + q1 + ")")
		}

	}
	if x.select3 != "" {
		q1 := ""
		if x.select3 == "docNumber" {
			q1 += ("id:(" + x.val3 + "* )")
		} else if x.select3 == "docTitle" {
			q1 += ("title:(" + strings.Replace(x.val3, " ", "*", -1) + "* )")
		} else if x.select3 == "docContent" {
			q1 += ("body:( *" + strings.Replace(x.val3, " ", "*", -1) + "* )")
		} else if x.select3 == "docUser" {
			q1 += (" creator:(" + strings.Replace(x.val3, " ", "*", -1) + ") approver:(" + strings.Replace(x.val3, " ", "*", -1) + ") authorizer:(" + strings.Replace(x.val3, " ", "*", -1) + ") initiator:(" + strings.Replace(x.val3, " ", "*", -1) + ") reviewer:(" + strings.Replace(x.val3, " ", "*", -1) + ")")
		}
		if query == "" {
			query += ("(" + q1 + ")")
		} else {
			query += (" AND (" + q1 + ")")
		}
	}
	return query
}
func makeDepartmentQuery(x formData) string {
	query := ""
	if x.d1 != "" {
		query += "docDepartment:D1"
	}
	if x.d2 != "" {
		query += " docDepartment:D2"
	}
	if x.d3 != "" {
		query += " docDepartment:D3"
	}
	return query
}
func makeProcessQuery(x formData) string {
	query := ""
	if x.everyone != "" {
		query += "docProcess:Everyone"
	}
	if x.onebyone != "" {
		query += " docProcess:OneByOne"
	}
	if x.anyone != "" {
		query += " docProcess:Anyone"
	}
	return query
}
func makeTypeQuery(x formData) string { //will return an empty string
	query := ""
	if x.sop != "" {
		query += "docType:SOP"
	}
	if x.hr != "" {
		query += " docType:HR"
	}
	if x.stp != "" {
		query += " docType:STP"
	}
	return query
}
func makedateQuery(x formData) string { //never returns an empty string
	bInit := "2000-01-01T00:00:00Z"
	bExp := "2000-01-01T00:00:00Z"
	bEff := "2000-01-01T00:00:00Z"
	eInit := "*"
	eEff := "*"
	eExp := "*"
	dQinit := "initTime:"
	dQeff := "effDate:"
	dQexp := "expDate:"
	if x.initFrom != "" {
		bInit = x.initFrom + "T00:00:00Z"
	}
	if x.initTo != "" {
		eInit = x.initTo + "T23:59:59Z"
	}
	if x.effFrom != "" {
		bEff = x.effFrom + "T00:00:00Z"
	}
	if x.effTo != "" {
		eEff = x.effTo + "T23:59:59Z"
	}
	if x.expFrom != "" {
		bExp = x.expFrom + "T00:00:00Z"
	}
	if x.expTo != "" {
		eExp = x.expTo + "T23:59:59Z"
	}
	return dQinit + "[" + bInit + " TO " + eInit + "]" + " AND " + dQeff + "[" + bEff + " TO " + eEff + "]" + " AND " + dQexp + "[" + bExp + " TO " + eExp + "]"
}

func printFormData(x formData) { //func used for debuging
	fmt.Println("sop :#" + x.sop + "#")
	fmt.Println("hr:#" + x.hr + "#")
	fmt.Println("stp:#" + x.stp + "#")
	fmt.Println("everyone:#" + x.everyone + "#")
	fmt.Println("onebyone:#" + x.onebyone + "#")
	fmt.Println("anyone:#" + x.anyone + "#")
	fmt.Println("d1:#" + x.d1 + "#")
	fmt.Println("d2:#" + x.d2 + "#")
	fmt.Println("d3:#" + x.d3 + "#")
	fmt.Println("initFrom:#" + x.initFrom + "#")
	fmt.Println("initTo:#" + x.initTo + "#")
	fmt.Println("effFrom:#" + x.effFrom + "#")
	fmt.Println("effTo:#" + x.effTo + "#")
	fmt.Println("expFrom:#" + x.expFrom + "#")
	fmt.Println("expTo:#" + x.expTo + "#")
	fmt.Println("select1:#" + x.select1 + "#")
	fmt.Println("val1:#" + x.val1 + "#")
	fmt.Println("select2:#" + x.select2 + "#")
	fmt.Println("val2:#" + x.val2 + "#")
	fmt.Println("select3:#" + x.select3 + "#")
	fmt.Println("val3:#" + x.val3 + "#")
}
func validate(x formData) bool {

	if !(x.sop == "" || x.sop == "on") {
		return false
	}
	if !(x.hr == "" || x.hr == "on") {
		return false
	}
	if !(x.stp == "" || x.stp == "on") {
		return false
	}
	if !(x.d1 == "" || x.d1 == "on") {
		return false
	}
	if !(x.d2 == "" || x.d2 == "on") {
		return false
	}
	if !(x.d3 == "" || x.d3 == "on") {
		return false
	}
	if !(x.everyone == "" || x.everyone == "on") {
		return false
	}
	if !(x.onebyone == "" || x.onebyone == "on") {
		return false
	}
	if !(x.anyone == "" || x.anyone == "on") {
		return false
	}
	isKeyword := regexp.MustCompile(`^[A-Za-z0-9 ]+$`).MatchString
	isDate := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`).MatchString
	isAlphaNumeric := regexp.MustCompile(`^[A-Za-z0-9]+$`).MatchString
	isallSpaces := regexp.MustCompile(`^[ ]+$`).MatchString
	if x.initFrom != "" {
		if isDatevalid(x.initFrom) != true || !isDate(x.initFrom) {
			return false
		}
	}
	if x.initTo != "" {
		if isDatevalid(x.initTo) != true || !isDate(x.initTo) {
			return false
		}
	}
	if x.effFrom != "" {
		if isDatevalid(x.effFrom) != true || !isDate(x.effFrom) {
			return false
		}
	}
	if x.effTo != "" {
		if isDatevalid(x.effTo) != true || !isDate(x.effTo) {
			return false
		}
	}
	if x.expFrom != "" {
		if isDatevalid(x.expFrom) != true || !isDate(x.expFrom) {
			return false
		}
	}
	if x.expTo != "" {
		if isDatevalid(x.expTo) != true || !isDate(x.expTo) {
			return false
		}
	}
	if x.select1 != "" {
		if x.select1 == "docNumber" {
			if !isAlphaNumeric(x.val1) || isallSpaces(x.val1) {
				return false
			}

		} else if x.select1 == "docContent" || x.select1 == "docUser" || x.select1 == "docTitle" {
			if !isKeyword(x.val1) || isallSpaces(x.val1) {
				return false
			}
		} else {
			return false
		}
	} else {
		if x.val1 != "" {
			return false
		}
	}
	if x.select2 != "" {
		if x.select2 == "docNumber" {
			if !isAlphaNumeric(x.val2) || isallSpaces(x.val2) {
				return false
			}
		} else if x.select2 == "docContent" || x.select2 == "docUser" || x.select2 == "docTitle" {
			if !isKeyword(x.val2) || isallSpaces(x.val2) {
				return false
			}
		} else {
			return false
		}
	} else {
		if x.val2 != "" {
			return false
		}
	}
	if x.select3 != "" {
		if x.select1 == "docNumber" {
			if !isAlphaNumeric(x.val3) || isallSpaces(x.val3) {
				return false
			}
		} else if x.select3 == "docContent" || x.select3 == "docUser" || x.select3 == "docTitle" {
			if !isKeyword(x.val3) || isallSpaces(x.val3) {
				return false
			}
		} else {
			return false
		}
	} else {
		if x.val3 != "" {
			return false
		}
	}
	return true

}
func isEmpty(x formData) bool {
	if (x.anyone + x.d1 + x.d2 + x.d3 + x.effFrom + x.effTo + x.everyone + x.expFrom + x.expTo + x.hr + x.initFrom + x.initTo + x.onebyone + x.select1 + x.select2 + x.select3 + x.sop + x.stp + x.val1 + x.val2 + x.val3) == "" {
		return false
	}
	return true
}
