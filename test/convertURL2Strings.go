package test

import (
	"fmt"
	"os"
)

//ConvertURL2Strings ... func
func ConvertURL2Strings(url string, tempFile string) ([]string, error) {

	data, err := ConvertURLtotext(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	f, err := os.Create(tempFile + ".pdf")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	f.Sync()

	paras := ConvertPDF2StringSlice(tempFile + ".pdf")

	for _, e := range paras {
		fmt.Println(e, "this is end of mylife")
	}
	err = os.Remove(tempFile + ".pdf")
	if err != nil {
		fmt.Println(err)
	}
	return paras, nil
}
