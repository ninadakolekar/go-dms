package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/buger/jsonparser"
)

func DBLPResponse() []paper {

	dbconf := []string{"VLDB", "SIGMOD", "PODS", "ICDE", "ICDT", "EDBT", "SIGKDD", "ICDM", "IJCAI", "AAAI", "ICML", "UAI", "UMAP", "NIPS", "AAMAS", "ACL", "AIED", "ITS", "SIGIR", "WWW", "ICIS", "PPoPP", "PACT", "IPDPS", "ICPP", "Euro-Par", "SIGGRAPH", "CVPR", "ICCV", "I3DG", "ACM-MM", "SIGCOMM", "PERFORMANCE", "SIGMETRICS", "INFOCOM", "MOBICOM", "IEEE", "CCS", "SOSP", "OSDI", "FOCS", "STOC", "ICALP", "SODA", "ISMB"}
	dbconf = []string{"VLDB", "EDBT", "UAI", "ACL"}
	dbconf = []string{"ICDE", "ICDM", "ICIS", "PACT", "IPDPS", "ICPP", "CVPR", "ICCV", "INFOCOM", "FOCS"}
	papers := []paper{}

	for _, conf := range dbconf {

		ch := make(chan []paper)

		go func(conf string) {

			ch <- getPapersDataset(conf)

		}(conf)

		listPaperG2 := <-ch
		papers = append(papers, listPaperG2...)
	}

	fmt.Println("TOTAL: ", len(papers))

	authors := getAuthors(papers)

	fmt.Println("TOTAL AUTHORS: ", len(authors))

	arr := []string{}
	for _, p := range papers {
		arr = append(arr, "\""+p.url+"\",")
	}

	fmt.Println(arr)

	return papers
}

// Paper model & related functions

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

const (
	minAuthors = 3
	minRepeats = 1
)

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

func getPapersDataset(conf string) []paper {

	listPaperG2 := []paper{}

	listPapers := getConfPapers(conf)
	listAuthorsG2 := getRepeatedAuthors(getConfAuthors(conf))

	for _, p := range listPapers {

		if len(p.authors) >= minAuthors {

			for _, a := range p.authors {

				if findPos(listAuthorsG2, a) != -1 {

					listPaperG2 = append(listPaperG2, p)
					break
				}
			}
		}

	}

	return listPaperG2
}

func getConfAuthors(conf string) []author {

	listAuthor := []author{}

	resp, err := getJSONResponse(conf)

	if err != nil {
		log.Fatal(err, " Line 212")
	}

	jsonparser.ArrayEach(resp, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {

		_, _, _, err = jsonparser.Get(value, "info", "ee")

		if err == nil {

			authors, err := toString(jsonparser.Get(value, "info", "authors", "author"))

			if err == nil {

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

func getRepeatedAuthors(authors []author) []author {

	listAuthorG2 := []author{}

	for _, e := range authors {
		if e.count >= minRepeats {
			listAuthorG2 = append(listAuthorG2, e)
		}
	}
	return listAuthorG2
}

func getAuthors(papers []paper) []string {

	author := []string{}

	for _, p := range papers {
		for _, a := range p.authors {
			if !contains(author, a) {
				author = append(author, a)
			}
		}
	}

	return author

}

// Convert paper to document

// func convert(papers []paper) []models.InactiveDoc {

// 	docs := []models.InactiveDoc{}

// 	for _, p := range papers {

// 		ch := make(chan models.InactiveDoc)

// 		go func(p paper, ch chan models.InactiveDoc) {
// 			ch <- convertToDoc(p)
// 		}(p, ch)

// 		document := <-ch
// 		docs = append(docs, document)

// 	}

// 	return docs
// }

// func convertToDoc(p paper) models.InactiveDoc {

// }

func getUniqueURL(papers []paper) []string {

	unique := []string{}

	for _, p := range papers {
		if !contains(unique, findDomain(p.url)) {
			unique = append(unique, findDomain(p.url))
		}
	}

	return unique
}

// Utility Functions
func findDomain(url string) string {

	url = strings.Split(url, "//")[1]
	//	url = strings.Split(url, "/")[0]

	return url

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

func contains(list []string, x string) bool {

	for _, item := range list {
		if item == x {
			return true
		}
	}

	return false
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
