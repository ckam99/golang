package sqlb

import (
	"fmt"
	"strings"
)

type QueryFilter struct {
	Limit, Offset int64
	OrderBy, Sort string
	Stmt          string
	args          []interface{}
	currentTag    string
}

func Query(baseQuery string) *QueryFilter {
	return &QueryFilter{
		Stmt: baseQuery,
		args: make([]interface{}, 0),
	}
}

func (q *QueryFilter) Where(column string, op string, value ...interface{}) *QueryFilter {
	q.clause("where", column, op, value...)
	return q
}

func (q *QueryFilter) Or(column string, op string, value ...interface{}) *QueryFilter {
	if q.currentTag == "where" || q.currentTag == "having" {
		switch op {
		case "in":
			q.in(column, "or", value...)
		case "between":
			q.between(column, "or", value[0], value[0])
		default:
			q.Stmt += fmt.Sprintf(" or %s %s $%d", column, op, len(q.args)+1)
			q.args = append(q.args, value[0])
		}
		return q
	}
	panic("function `Or` should be called after where or having statement")
}

func (q *QueryFilter) GroupBy(columns ...string) *QueryFilter {
	s := fmt.Sprintf(" group by %s", strings.Join(columns, ","))
	if strings.Contains(strings.ToLower(q.Stmt), "group by") {
		q.Stmt = strings.ReplaceAll(strings.ToLower(q.Stmt), "group by", s+",")
	} else {
		q.Stmt += s
	}
	q.currentTag = "group by"
	return q
}

func (q *QueryFilter) Having(column string, op string, value ...interface{}) *QueryFilter {
	if q.currentTag == "having" || q.currentTag == "group by" || strings.Contains(strings.ToLower(q.Stmt), "group by"){
		q.clause("having", column, op, value...)
		return q
	}
	panic("function `Having` should be called after group by statement")
}

func (q *QueryFilter) Build() []interface{} {

	if q.OrderBy != "" {
		q.Stmt += fmt.Sprintf(" order by %s", q.OrderBy)
		if strings.ToLower(q.Sort) == "desc" {
			q.Stmt += " " + q.Sort
		}
	}

	if q.Limit > 0 {
		q.Stmt += fmt.Sprintf(` limit %d`, q.Limit)
	}
	if q.Offset > 0 {
		q.Stmt += fmt.Sprintf(` offset %d;`, q.Offset)
	}
	return q.args
}

func (q *QueryFilter) clause(clause string, column string, op string, value ...interface{}) *QueryFilter {
	if strings.Count(q.Stmt, clause) > 0 {
		q.Stmt += " and"
	} else {
		q.Stmt += " " + clause
	}
	switch op {
	case "in":
		q.in(column, "", value...)
	case "between":
		q.between(column, "", value[0], value[0])
	default:
		q.Stmt += fmt.Sprintf(" %s %s $%d", column, op, len(q.args)+1)
		q.args = append(q.args, value[0])
	}
	q.currentTag = clause
	return q
}

func (q *QueryFilter) between(column string, logic string, from interface{}, to interface{}) {
	if strings.Count(q.Stmt, q.currentTag) > 0 {
		q.Stmt += " " + logic
	} else {
		q.Stmt += " " + q.currentTag
	}
	q.Stmt += fmt.Sprintf(" %s between $%d and $%d", column, len(q.args)+1, len(q.args)+2)
	q.args = append(q.args, from, to)
}

func (q *QueryFilter) in(column string, logic string, values ...any) {
	if strings.Count(q.Stmt, q.currentTag) > 0 {
		q.Stmt += fmt.Sprintf(" %s %s in (", logic, column)
	} else {
		q.Stmt += fmt.Sprintf(" %s %s in (", q.currentTag, column)
	}

	for k, v := range values {
		q.Stmt += fmt.Sprintf("$%d", len(q.args)+1)
		if k < len(values)-1 {
			q.Stmt += ","
		}
		q.args = append(q.args, v)
	}
	q.Stmt += ")"
}
