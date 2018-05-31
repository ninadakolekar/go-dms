package test

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/unidoc/unidoc/pdf/extractor"
	pdf "github.com/unidoc/unidoc/pdf/model"
)

//ConvertPDF2StringSlice ... converts pdf to string
func ConvertPDF2StringSlice(url string) {
	fmt.Println(url)

	inputPath := "test/sample.pdf"
	pages, err := outputPdfText(inputPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	//i := len(pages)
	for i, p := range pages {
		fmt.Println("\n\n\n", "pagenumber:", i)
		fmt.Println(p)
	}
	// fmt.Println(i)
}
func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// outputPdfText prints out contents of PDF file to stdout.
func outputPdfText(inputPath string) ([]string, error) {
	re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	pages := []string{}

	f, err := os.Open(inputPath)
	if err != nil {
		return pages, err
	}

	defer f.Close()

	pdfReader, err := pdf.NewPdfReader(f)
	if err != nil {
		return pages, err
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return pages, err
	}

	// fmt.Printf("--------------------\n")
	// fmt.Printf("PDF to text extraction:\n")
	// fmt.Printf("--------------------\n")
	for i := 0; i < numPages; i++ {
		pageNum := i + 1

		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			return pages, err
		}

		ex, err := extractor.New(page)
		if err != nil {
			return pages, err
		}

		text, err := ex.ExtractText()
		if err != nil {
			return pages, err
		}
		text = re_leadclose_whtsp.ReplaceAllString(text, "")
		text = re_inside_whtsp.ReplaceAllString(text, " ")
		pages = append(pages, standardizeSpaces(text))
	}

	return pages, nil
}
