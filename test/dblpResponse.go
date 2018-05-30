package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/buger/jsonparser"
)

type authour struct {
	name  string
	count int
}

func findPos(s []authour, a string) int {
	for i, e := range s {
		if e.name == a {
			return i
		}
	}
	return -1
}

func DBLPResponse() {

	dbconf := []string{"VLDB", "SIGMOD", "PODS", "ICDE", "ICDT", "EDBT"}
	var A []authour
	for _, conf := range dbconf {

		resp, err := getJSONResponse(conf)

		if err != nil {
			log.Fatal(err, " Line 21")
		}
		fmt.Println(conf)
		counter := 0
		jsonparser.ArrayEach(resp, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {

			_, _, _, errr := jsonparser.Get(value, "info", "ee")

			if errr != nil {
				// log.Println(err)
				fmt.Println(err)
			} else {
				b, _, _, _ := jsonparser.Get(value, "info", "authors", "author")
				bx := string(b)
				bxx := []string{}
				i := strings.Index(bx, "\"")
				for true {
					i = strings.Index(bx, "\"")
					if i <= -1 {
						break
					}
					bx = bx[:i] + "*" + bx[i+1:]

					first := i
					i = strings.Index(bx, "\"")
					last := i

					if last > -1 {
						bx = bx[:i] + "*" + bx[i+1:]
						bxx = append(bxx, bx[first+1:last])
					}
				}
				for _, e := range bxx {
					i := findPos(A, e)
					if i == -1 {
						var xx authour
						xx.name = e
						xx.count = 0
						A = append(A, xx)
					} else {
						A[i].count++
					}
				}
				counter++

			}

		}, "result", "hits", "hit")

		if err != nil {
			log.Fatal(err)
		}
	}
	for i, e := range A {
		fmt.Println("authour S.no :", i, " name:", e.name, " repeatation:", e.count)
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
