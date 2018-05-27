package utility

// StringInSlice ... Checks membership of given string in given string slice
func StringInSlice(a string, list []string) bool {

	if list == nil {
		return false
	}

	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
