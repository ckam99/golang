package sqlx

import (
	"fmt"
	"reflect"
	"strings"
)

func normalizeInsertValueQuery(query string, fields ...interface{}) string {
	q := []string{}
	for k := range fields {
		q = append(q, fmt.Sprint("$", k+1))
	}
	return query + fmt.Sprintf(" VALUES(%s)", strings.Join(q, ","))
}

func normalizeWhereQuery(query, op, field, condition string, value interface{}) string {
	if strings.Contains(query, "WHERE") {
		query += " " + op + " "
	} else {
		query += " WHERE "
	}
	query += fmt.Sprintf("%s%s%v", field, condition, value)
	return query
}

func normalizeWhereInQuery(query, op, field, condition string, values interface{}) string {
	if strings.Contains(query, "WHERE") {
		query += " " + op + " "
	} else {
		query += " WHERE "
	}
	query += fmt.Sprintf("%s %s (%s)", field, condition, parseIntf(values))
	return query
}

func parseIntf(slice interface{}) string {
	slc := reflect.ValueOf(slice)
	if slc.Kind() != reflect.Slice {
		return ""
	}
	if slc.IsNil() {
		return ""
	}
	ret := make([]string, slc.Len())
	for i := 0; i < slc.Len(); i++ {
		ret[i] = fmt.Sprint(slc.Index(i))
	}
	return strings.Join(ret, ",")
}
