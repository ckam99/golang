package sqlb

import (
	"fmt"
	"strings"
)

func CleanSQL(query string, args ...interface{}) string {
	s := query
	for k, v := range args {
		s = strings.Replace(s, fmt.Sprintf("$%d", k+1), fmt.Sprint(v), -1)
	}
	return strings.ReplaceAll(strings.Trim(strings.ReplaceAll(
		strings.ReplaceAll(query, "\t", ""), "\n", " "), " "), "  ", " ")
}
