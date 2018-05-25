package controllers

import (
	"fmt"
	"html"
	"net/http"
	"regexp"

	constant "github.com/ninadakolekar/aizant-dms/src/constants"
	solr "github.com/rtt/Go-Solr"
)

//ProcessDocSearch ... process new search
func ProcessDocSearch(w http.ResponseWriter, r *http.Request) {

	links := []lINK{}
	sortOrder := "default"
	query := buildQuery(r)
	fmt.Println("query:", query)

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
		fmt.Fprintf(w, "<h1>no results found</h1>")
	} else {
		for i := 0; i < results.Len(); i++ {
			links = append(links, convertTolINK(results.Get(i)))
		}

		links = sortby(links, sortOrder)
		fmt.Println("after sorting \n", links) //Debug
		s1 := "<li class='collection-item avatar'><i class='material-icons circle green'>insert_drive_file</i><span class='title'>"
		s2 := "</span><p>"
		s3 := "</p><a href='"
		s4 := "' class = 'secondary-content'><br><i class='material-icons'>send</i></a></li>"

		ret := ""
		for _, e := range links {
			ret += (s1 + e.DocName + s2 + "intiated:" + e.Idate + s3 + "/doc/view/" + e.DocId + s4)
		}

		fmt.Fprintf(w, ret)
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

func buildQuery(r *http.Request) string {

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

	printFormData(information)

	query := makedateQuery(information)
	return query
}

func makedateQuery(x formData) string {
	isDate := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`).MatchString
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
		fmt.Println(x.initFrom, "special attention : isvalid ", isDatevalid(x.initFrom), isDate(x.initFrom))
		if isDatevalid(x.initFrom) == true || isDate(x.initFrom) {
			bInit = x.initFrom + "T00:00:00Z"
		}
	}
	if x.initTo != "" {
		if isDatevalid(x.initTo) == true || isDate(x.initTo) {
			eInit = x.initTo + "T23:59:59Z"
		}
	}
	if x.effFrom != "" {
		if isDatevalid(x.effFrom) == true || isDate(x.effFrom) {
			bEff = x.effFrom + "T00:00:00Z"
		}
	}
	if x.effTo != "" {
		if isDatevalid(x.effTo) == true || isDate(x.effTo) {
			eEff = x.effTo + "T23:59:59Z"
		}
	}
	if x.expFrom != "" {
		if isDatevalid(x.expFrom) == true || isDate(x.expFrom) {
			bExp = x.expFrom + "T00:00:00Z"
		}
	}
	if x.expTo != "" {
		if isDatevalid(x.expTo) == true || isDate(x.expTo) {
			eExp = x.expTo + "T23:59:59Z"
		}
	}
	return dQinit + "[" + bInit + " TO " + eInit + "]" + " AND " + dQeff + "[" + bEff + " TO " + eEff + "]" + " AND " + dQexp + "[" + bExp + " TO " + eExp + "]"
}

func printFormData(x formData) {
	fmt.Println("sop :", x.sop)          //Debug
	fmt.Println("hr:", x.hr)             //Debug
	fmt.Println("stp:", x.stp)           //Debug
	fmt.Println("everyone:", x.everyone) //Debug
	fmt.Println("onebyone:", x.onebyone) //Debug
	fmt.Println("anyone:", x.anyone)     //Debug
	fmt.Println("d1:", x.d1)             //Debug
	fmt.Println("d2:", x.d2)             //Debug
	fmt.Println("d3:", x.d3)             //Debug
	fmt.Println("initFrom:", x.initFrom) //Debug
	fmt.Println("initTo:", x.initTo)     //Debug
	fmt.Println("effFrom:", x.effFrom)   //Debug
	fmt.Println("effTo:", x.effTo)       //Debug
	fmt.Println("expFrom:", x.expFrom)   //Debug
	fmt.Println("expTo:", x.expTo)       //Debug
	fmt.Println("select1:", x.select1)   //Debug
	fmt.Println("val1:", x.val1)         //Debug
	fmt.Println("select2:", x.select2)   //Debug
	fmt.Println("val2:", x.val2)         //Debug
	fmt.Println("select3:", x.select3)   //Debug
	fmt.Println("val3:", x.val3)         //Debug
}
