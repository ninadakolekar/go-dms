package main

import (
	"fmt"
	"net/http"

	constants "github.com/ninadakolekar/go-dms/src/constants"
	router "github.com/ninadakolekar/go-dms/src/routes"
)

func main() {

	// test.DBLPResponse()
	// test.InsertUsers(papers)
	// test.InsertPDF(papers)
	//http://www.vldb.org/conf/2007/papers/demo/p1422-antova.pdf    //vldb
	//https://ieeexplore.ieee.org/ielx7/7929494/7929895/07929967.pdf     // ieeee
	// text, err := test.ConvertURL2Strings("https://ieeexplore.ieee.org/ielx7/7929494/7929895/07929967.pdf", "temp")
	urls := []string{"https://ieeexplore.ieee.org/ielx7/7929494/7929895/07929967.pdf", "https://ieeexplore.ieee.org/ielx7/7929494/7929895/07930060.pdf", "https://ieeexplore.ieee.org/ielx7/7929494/7929895/07930084.pdf", "https://ieeexplore.ieee.org/ielx7/7929494/7929895/07929945.pdf", "https://ieeexplore.ieee.org/ielx7/7929494/7929895/07930014.pdf", "https://ieeexplore.ieee.org/ielx7/7929494/7929895/07930044.pdf", "https://ieeexplore.ieee.org/ielx7/7929494/7929895/07930028.pdf", "https://ieeexplore.ieee.org/ielx7/7929494/7929895/07929952.pdf", "https://ieeexplore.ieee.org/ielx7/7929494/7929895/07930050.pdf", "https://ieeexplore.ieee.org/ielx7/7929494/7929895/07929947.pdf", "https://ieeexplore.ieee.org/ielx7/7929494/7929895/07929912.pdf", "https://ieeexplore.ieee.org/ielx7/7929494/7929895/07930042.pdf", "https://ieeexplore.ieee.org/ielx7/7929494/7929895/07930031.pdf", "https://ieeexplore.ieee.org/ielx7/7929494/7929895/07930009.pdf", "https://ieeexplore.ieee.org/ielx7/7929494/7929895/07930071.pdf", "https://ieeexplore.ieee.org/ielx7/7929494/7929895/07930100.pdf", "https://ieeexplore.ieee.org/ielx7/7929494/7929895/07929929.pdf", "https://ieeexplore.ieee.org/ielx7/7929494/7929895/07930089.pdf", "https://ieeexplore.ieee.org/ielx7/7929494/7929895/07930062.pdf", "https://ieeexplore.ieee.org/ielx7/7929494/7929895/07929916.pdf", "https://ieeexplore.ieee.org/ielx7/7929494/7929895/07929951.pdf", "https://ieeexplore.ieee.org/ielx7/7929494/7929895/07929927.pdf", "https://ieeexplore.ieee.org/ielx7/7929494/7929895/07930037.pdf"}
	fmt.Println(len(urls))
	/*
		for i, e := range urls {
			ch := make(chan int)
			go func(ch chan int, e string, i int) {
				text, err := test.ConvertURL2Strings(e, strconv.Itoa(i))
				if err != nil {
					fmt.Println("ERROR app.go Line 19: ", err)
					ch <- 0
				} else {
					fmt.Println(text[0])
					ch <- 1
				}
			}(ch, e, i)

			fmt.Println(<-ch)

		}
	*/
	// text, err := test.ConvertURL2Strings("https://ieeexplore.ieee.org/ielx7/7929494/7929895/07929967.pdf", "temp12")
	// if err != nil {
	// 	fmt.Println("ERROR app.go Line 19: ", err)
	// 	//ch <- 0
	// } else {
	// 	fmt.Println(text[0])
	// 	//ch <- 1
	// }

	r := router.GetRouter()
	http.ListenAndServe(constants.ApplicationPort, r)
}
