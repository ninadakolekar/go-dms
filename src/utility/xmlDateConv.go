package utility

// XMLDate ... Get Date in XML format
func XMLDate(date string) string {
	return date + "T00:00:01Z"
}
