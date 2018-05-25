package controllers

import (
	"fmt"
	"net/http"
)

//ProcessDocSearch ... process new search
func ProcessDocSearch(w http.ResponseWriter, r *http.Request) {
	buildQuery(r)
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
		sop:      r.FormValue("SOP"),
		hr:       r.FormValue("HR"),
		stp:      r.FormValue("STP"),
		everyone: r.FormValue("Everyone"),
		onebyone: r.FormValue("OneByOne"),
		anyone:   r.FormValue("Anyone"),
		d1:       r.FormValue("D1"),
		d2:       r.FormValue("D2"),
		d3:       r.FormValue("D3"),
		initFrom: r.FormValue("intiFrom"),
		initTo:   r.FormValue("initTo"),
		effFrom:  r.FormValue("effFrom"),
		effTo:    r.FormValue("effTo"),
		expFrom:  r.FormValue("expFrom"),
		expTo:    r.FormValue("expTo"),
		select1:  r.FormValue("select1"),
		val1:     r.FormValue("val1"),
		select2:  r.FormValue("select2"),
		val2:     r.FormValue("val2"),
		select3:  r.FormValue("select3"),
		val3:     r.FormValue("val3"),
	}

	printFormData(information)

	return "*"
}

func makedateQuery(s string, e string) string {

	return ""
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
