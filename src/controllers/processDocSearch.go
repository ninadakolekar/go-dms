package controllers

import (
	"fmt"
	"html"
	"html/template"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"

	constant "github.com/ninadakolekar/aizant-dms/src/constants"
	solr "github.com/rtt/Go-Solr"
)

type lINK struct {
	DocName string
	Idate   string
	DocId   string
}
type docNameSorter []lINK
type idateSorter []lINK

func (a docNameSorter) Len() int           { return len(a) }
func (a docNameSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a docNameSorter) Less(i, j int) bool { return a[i].DocName < a[j].DocName }

func (a idateSorter) Len() int           { return len(a) }
func (a idateSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a idateSorter) Less(i, j int) bool { return a[i].Idate < a[j].Idate }

//ProcessDocSearch ... process doc search
func ProcessDocSearch(w http.ResponseWriter, r *http.Request) {

	alert := false
	data := false
	alertmsg := "no msg"
	links := []lINK{}
	if r.Method == "POST" {
		r.ParseForm()
		sCriteria := []string{"*", "*", "*", "*", "*", "*", "*"}
		sKeyword := []string{"*", "*", "*", "*", "*", "*", "*"}
		sortOrder := html.EscapeString(r.Form["sort"][0])

		for i := 0; i < 6; i++ {
			if len(r.Form["criteria"+strconv.Itoa(i+1)]) > 0 {
				sCriteria[i] = html.EscapeString(r.Form["criteria"+strconv.Itoa(i+1)][0])
				sKeyword[i] = html.EscapeString(r.Form["searchKeyword"+strconv.Itoa(i+1)][0])
			}
		}
		for i := 0; i < 6; i++ {
			sKeyword[i] = removeIntialEndingspaces(sKeyword[i])
		}

		if validateSearchForm(sCriteria, sKeyword) == true {

			query := makeSearchQuery(sCriteria, sKeyword)

			s, err := solr.Init(constant.SolrHost, constant.SolrPort, constant.DocsCore)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(query) //Debug
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
				alert = true
				alertmsg = "No Results Found!"
			} else {
				for i := 0; i < results.Len(); i++ {
					links = append(links, lINK{results.Get(i).Field("title").(string), results.Get(i).Field("initTime").(string), results.Get(i).Field("id").(string)})
				}
				data = true

				links = sortby(links, sortOrder)
				//	fmt.Println("after sorting \n", links) //Debug
			}

		} else {
			alert = true
			alertmsg = "Invalid Search Query!"
		}
	}

	tmpl := template.Must(template.ParseFiles("templates/searchDoc.html"))
	tmpl.Execute(w, struct {
		Alertb   bool
		Alertmsg string
		Datab    bool
		Data     []lINK
	}{alert, alertmsg, data, links})
}

func sortby(l []lINK, so string) []lINK {

	if so == "alexical" {
		sort.Sort(docNameSorter(l))

	} else if so == "aTime" {
		sort.Sort(idateSorter(l))

	}
	return l
}

func makeSearchQuery(sC []string, sK []string) string {
	validCriterion := []string{"docNumber", "docName", "docKeyword", "initiator", "creator", "reviewer", "approver", "auth", "dept", "from Init Date", "from Eff Date", "from Exp Date", "till Init Date", "till Eff Date", "till Exp Date"}
	validQueryPrifex := []string{"id:", "title:", "body:", "initiator:", "creator:", "reviewer:", "approver:", "authorizer:", "docDepartment:", "initTime:", "effDate:", "expDate:", "initTime:", "effDate:", "expDate:"}

	querys := []string{}
	counter := 0
	for i, sc := range sC {
		for j, v := range validCriterion {
			if v == sc {

				if j == 10 || j == 9 || j == 11 {
					querys = append(querys, validQueryPrifex[j]+"["+sK[i]+"T23:59:59Z TO *]")
				} else if j == 12 || j == 13 || j == 14 {
					querys = append(querys, validQueryPrifex[j]+"[ 2000-01-01T00:00:58Z TO "+sK[i]+"T23:59:59Z ]")
				} else if j == 0 || j == 1 {
					querys = append(querys, validQueryPrifex[j]+sK[i]+"*")
				} else if j == 2 {
					strs := strings.Split(sK[i], " ")
					str := "*"
					for _, e := range strs {
						str = str + e + "*"
					}
					querys = append(querys, validQueryPrifex[j]+str)
				} else {
					querys = append(querys, validQueryPrifex[j]+sK[i])
				}

				// fmt.Println(querys[counter])
				counter++
			}
		}
	}
	query := ""
	for i, q := range querys {
		if i == 0 {
			query = q
		} else {
			query += (" AND " + q)
		}
	}

	return query
}
func isDatevalid(s string) bool {
	y, err := strconv.Atoi(s[0:4])
	if err != nil {
		return false
	}

	m, err := strconv.Atoi(s[5:7])
	if err != nil {
		return false
	}
	d, err := strconv.Atoi(s[8:10])
	if err != nil {
		return false
	}

	if (y%4 == 0 && y%100 == 0) || (y%4 != 0) {
		if m == 1 || m == 3 || m == 5 || m == 7 || m == 8 || m == 10 || m == 12 {
			if d < 0 || d > 31 {
				return false
			}
		} else if m == 4 || m == 6 || m == 9 || m == 11 {
			if d < 0 || d > 30 {
				return false
			}
		} else if m == 2 {
			if d < 0 || d > 29 {
				return false
			}
		} else {
			return false
		}
	} else {
		if m == 1 || m == 3 || m == 5 || m == 7 || m == 8 || m == 10 || m == 12 {
			if d < 0 || d > 31 {
				return false
			}
		} else if m == 4 || m == 6 || m == 9 || m == 11 {
			if d < 0 || d > 30 {
				return false
			}
		} else if m == 2 {
			if d < 0 || d > 28 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
func validateSearchForm(sC []string, sK []string) bool {
	validCriterion := []string{"docNumber", "docName", "docKeyword", "initiator", "creator", "reviewer", "approver", "auth", "dept", "from Init Date", "from Eff Date", "from Exp Date", "till Init Date", "till Eff Date", "till Exp Date"}
	isKeyword := regexp.MustCompile(`^[A-Za-z0-9 ]+$`).MatchString
	isDate := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`).MatchString
	isAlphaNumeric := regexp.MustCompile(`^[A-Za-z0-9_ ]+$`).MatchString

	for j, sc := range sC {

		for i, v := range validCriterion {

			if sc == v {

				if i == 9 || i == 10 || i == 11 || i == 12 || i == 13 || i == 14 {
					if isDatevalid(sK[j]) == false || !isDate(sK[j]) {
						return false
					}
				} else if i == 2 {
					if isKeyword(sK[j]) == false {
						return false
					}
				} else {
					if isAlphaNumeric(sK[j]) == false {
						return false
					}
				}

			}
		}
	}
	return true
}
func removeIntialEndingspaces(str string) string {
	s := 0
	e := len(str) - 1

	for ; str[s] == ' '; s++ {

	}
	for e = len(str) - 1; str[e] == ' '; e-- {

	}

	return str[s : e+1]
}
