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
func ConvertPDF2StringSlice(inputPath string) []string {

	fmt.Println(inputPath)

	pages, err := outputPdfText(inputPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	return pages
}
func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

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
