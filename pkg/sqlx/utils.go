package sqlx

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func buildFilterQuery(fv []whereQuery, args *[]interface{}) string {
	var query strings.Builder
	query.WriteString(" where ")
	for k, v := range fv {
		query.WriteString(v.Column)
		if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
			query.WriteString(" in (")
			values := reflect.ValueOf(v.Value)
			for j := 0; j < values.Len(); j++ {
				query.WriteString(fmt.Sprintf("$%d", len(*args)+1))
				*args = append(*args, values.Index(j))
				if j < values.Len()-1 {
					query.WriteString(",")
				}
			}
			query.WriteString(")")
		} else {
			query.WriteString(fmt.Sprintf(" %s $%d", v.Operator, len(*args)+1))
			*args = append(*args, v.Value)
		}
		if k < len(fv)-1 {
			query.WriteString(" " + v.Type + " ")
		}
	}
	return query.String()
}

func ConvertStructToMap(v *interface{}) map[string]interface{} {
	m, _ := json.Marshal(v)
	var x map[string]interface{}
	_ = json.Unmarshal(m, &x)
	return x
}
