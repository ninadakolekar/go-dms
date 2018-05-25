package controllers

import (
	"fmt"
	"net/http"
)

//ProcessDocSearch ... process new search
func ProcessDocSearch(w http.ResponseWriter, r *http.Request) {
	buildQuery(r)
}
func buildQuery(r *http.Request) string {

	sop := r.FormValue("SOP")
	hr := r.FormValue("HR")
	stp := r.FormValue("STP")
	everyone := r.FormValue("Everyone")
	onebyone := r.FormValue("OneByOne")
	anyone := r.FormValue("Anyone")
	d1 := r.FormValue("D1")
	d2 := r.FormValue("D2")
	d3 := r.FormValue("D3")
	initFrom := r.FormValue("intiFrom")
	initTo := r.FormValue("initTo")
	effFrom := r.FormValue("effFrom")
	effTo := r.FormValue("effTo")
	expFrom := r.FormValue("expFrom")
	expTo := r.FormValue("expTo")
	select1 := r.FormValue("select1")
	val1 := r.FormValue("val1")
	select2 := r.FormValue("select2")
	val2 := r.FormValue("val2")
	select3 := r.FormValue("select3")
	val3 := r.FormValue("val3")
	fmt.Println("sop :", sop)          //Debug
	fmt.Println("hr:", hr)             //Debug
	fmt.Println("stp:", stp)           //Debug
	fmt.Println("everyone:", everyone) //Debug
	fmt.Println("onebyone:", onebyone) //Debug
	fmt.Println("anyone:", anyone)     //Debug
	fmt.Println("d1:", d1)             //Debug
	fmt.Println("d2:", d2)             //Debug
	fmt.Println("d3:", d3)             //Debug
	fmt.Println("initFrom:", initFrom) //Debug
	fmt.Println("initTo:", initTo)     //Debug
	fmt.Println("effFrom:", effFrom)   //Debug
	fmt.Println("effTo:", effTo)       //Debug
	fmt.Println("expFrom:", expFrom)   //Debug
	fmt.Println("expTo:", expTo)       //Debug
	fmt.Println("select1:", select1)   //Debug
	fmt.Println("val1:", val1)         //Debug
	fmt.Println("select2:", select2)   //Debug
	fmt.Println("val2:", val2)         //Debug
	fmt.Println("select3:", select3)   //Debug
	fmt.Println("val3:", val3)         //Debug

	return "*"
}
