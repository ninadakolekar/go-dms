package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/buger/jsonparser"
)

type author struct {
	name  string
	count int
}
type paper struct {
	id       string
	title    string
	authors  []string
	publDate string
	url      string
	venue    string
}

func findPos(s []author, a string) int {
	for i, e := range s {
		if e.name == a {
			return i
		}
	}
	return -1
}
func printPaperDetails(a paper) {
	fmt.Println("id :", a.id)
	fmt.Println("title :", a.title)
	fmt.Println("authors :", a.authors)
	fmt.Println("url :", a.url)
	fmt.Println("conference :", a.venue)
}
func DBLPResponse() {

	dbconf := []string{"VLDB", "SIGMOD", "PODS", "ICDE", "ICDT", "EDBT"}
	dbconf = []string{"SIGKDD", "ICDM"}
	dbconf = []string{"IJCAI", "AAAI", "ICML", "UAI", "UMAP"}
	dbconf = []string{"IJCAI"}
	for _, conf := range dbconf {
		listPaperG2 := getPapersG2(conf)
		for i, e := range listPaperG2 {
			fmt.Println("------------------------\n", i)
			printPaperDetails(e)
		}
	}
}

func getJSONResponse(conf string) ([]byte, error) {

	url := getVenueURL(conf)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func getVenueURL(conf string) string {
	return "http://dblp.org/search/publ/api?q=venue%3A" + conf + "%3A&format=json&h%3A1000&h=150"
}

func toString(value []byte, dataType jsonparser.ValueType, offset int, err error) (string, error) {
	if err != nil {
		return "", err
	}
	return string(value), err
}

func getConfPapers(conf string) []paper {

	papers := []paper{}

	resp, err := getJSONResponse(conf)

	if err != nil {
		log.Fatal(err, " Line 144")
	}

	jsonparser.ArrayEach(resp, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {

		var p paper

		skip := false

		p.url, err = toString(jsonparser.Get(value, "info", "ee"))
		if err != nil {
			skip = true
		}

		p.title, err = toString(jsonparser.Get(value, "info", "title"))
		if err != nil {
			skip = true
		}

		p.publDate, err = toString(jsonparser.Get(value, "info", "year"))
		if err != nil {
			skip = true
		}

		p.venue, err = toString(jsonparser.Get(value, "info", "venue"))
		if err != nil {
			skip = true
		}

		authors, err := toString(jsonparser.Get(value, "info", "authors", "author"))
		if err != nil {
			skip = true
		}

		p.id, err = toString(jsonparser.Get(value, "@id"))
		if err != nil {
			skip = true
		}

		if !skip {

			for _, author := range strings.Split(authors[1:len(authors)-1], ",") {
				p.authors = append(p.authors, author[1:len(author)-1])
			}

			papers = append(papers, p)

		}

	}, "result", "hits", "hit")

	return papers
}

func getPapersG2(conf string) []paper {
	listPaperG2 := []paper{}
	listPapers := getConfPapers(conf)
	listAuthorsG2 := getAuthorG2(getAuthors(conf))

	for _, papeR := range listPapers {

		flag := 0
		for _, authoR := range papeR.authors {
			if findPos(listAuthorsG2, authoR) != -1 {
				flag = 1
				break
			}
		}

		if flag != 0 {
			listPaperG2 = append(listPaperG2, papeR)
		}
	}

	return listPaperG2
}
func getAuthors(conf string) []author {

	listAuthor := []author{}

	resp, err := getJSONResponse(conf)

	if err != nil {
		log.Fatal(err, " Line 212")
	}

	jsonparser.ArrayEach(resp, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {

		_, _, _, err = jsonparser.Get(value, "info", "ee")

		if err != nil {

			fmt.Println("Line no 222", err) //Debug

		} else {

			authors, err := toString(jsonparser.Get(value, "info", "authors", "author"))

			if err != nil {
				fmt.Println("Line no 230 , empty authors", err)
			} else {

				for _, authorEach := range strings.Split(authors[1:len(authors)-1], ",") {

					index := findPos(listAuthor, authorEach[1:len(authorEach)-1])
					if index == -1 {
						listAuthor = append(listAuthor, author{name: authorEach[1 : len(authorEach)-1], count: 0})
					} else {
						listAuthor[index].count++
					}

				}
			}
		}

	}, "result", "hits", "hit")

	return listAuthor
}

func getAuthorG2(authors []author) []author {

	listAuthorG2 := []author{}
	for _, e := range authors {
		if e.count >= 2 {
			listAuthorG2 = append(listAuthorG2, e)
		}
	}
	return listAuthorG2
}
