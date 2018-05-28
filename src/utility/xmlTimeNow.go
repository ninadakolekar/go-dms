package utility

import (
	"strings"
	"time"
)

// XMLTimeNow ... Get current time in XML format
func XMLTimeNow() string {
	t := time.Now().UTC().String()
	x := strings.Fields(t)
	xmlTime := x[0] + "T" + x[1] + "Z"
	return xmlTime
}
