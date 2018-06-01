package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//ConvertURLtotext ... prints things in url
func ConvertURLtotext(url string) ([]byte, error) {

	fmt.Printf("HTML code of %s ...\n", url)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return html, nil
}
