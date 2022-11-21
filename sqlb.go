package sqlb

import (
	"fmt"
	"strings"
)

type QueryBuilder struct {
	Limit, Offset int64
	OrderBy, Sort string
	Stmt          string
	args          []interface{}
	currentTag    string
}

func Query(baseQuery string) *QueryBuilder {
	return &QueryBuilder{
		Stmt: baseQuery,
		args: make([]interface{}, 0),
	}
}

func (q *QueryBuilder) Build() *QueryBuilder {
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
	return q
}

func (q *QueryBuilder) Args() []interface{} {
	return q.args
}

func (q *QueryBuilder) Where(column string, op string, value ...interface{}) *QueryBuilder {
	q.clause("where", column, op, value...)
	return q
}

func (q *QueryBuilder) Or(column string, op string, value ...interface{}) *QueryBuilder {
	if q.currentTag == "where" || q.currentTag == "having" {
		switch op {
		case "in":
			q.in(column, "or", value...)
		case "is null":
			q.isNUll(column, "or")
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

func (q *QueryBuilder) GroupBy(columns ...string) *QueryBuilder {
	s := fmt.Sprintf(" group by %s", strings.Join(columns, ","))
	if strings.Contains(strings.ToLower(q.Stmt), "group by") {
		q.Stmt = strings.ReplaceAll(strings.ToLower(q.Stmt), "group by", s+",")
	} else {
		q.Stmt += s
	}
	q.currentTag = "group by"
	return q
}

func (q *QueryBuilder) Having(column string, op string, value ...interface{}) *QueryBuilder {
	if q.currentTag == "having" || q.currentTag == "group by" || strings.Contains(strings.ToLower(q.Stmt), "group by") {
		q.clause("having", column, op, value...)
		return q
	}
	panic("function `Having` should be called after group by statement")
}

func (q *QueryBuilder) clause(clause string, column string, op string, value ...interface{}) *QueryBuilder {
	if strings.Count(q.Stmt, clause) > 0 {
		q.Stmt += " and"
	} else {
		q.Stmt += " " + clause
	}

	switch op {
	case "in":
		q.in(column, "", value...)
	case "is null":
		q.isNUll(column, "")
	case "between":
		q.between(column, "", value[0], value[0])
	default:
		q.Stmt += fmt.Sprintf(" %s %s $%d", column, op, len(q.args)+1)
		q.args = append(q.args, value[0])
	}

	q.currentTag = clause
	return q
}

func (q *QueryBuilder) between(column string, logic string, from interface{}, to interface{}) {
	if strings.Count(q.Stmt, q.currentTag) > 0 {
		q.Stmt += " " + logic
	} else {
		q.Stmt += " " + q.currentTag
	}
	q.Stmt += fmt.Sprintf(" %s between $%d and $%d", column, len(q.args)+1, len(q.args)+2)
	q.args = append(q.args, from, to)
}

func (q *QueryBuilder) in(column string, logic string, values ...any) {
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

func (q *QueryBuilder) isNUll(column string, logic string) {
	if strings.Count(q.Stmt, q.currentTag) > 0 {
		q.Stmt += fmt.Sprintf(" %s %s is null", logic, column)
	} else {
		q.Stmt += fmt.Sprintf(" %s %s is null", q.currentTag, column)
	}
}

func CleanSQL(query string) string {
	return strings.ReplaceAll(strings.Trim(strings.ReplaceAll(
		strings.ReplaceAll(query, "\t", ""), "\n", " "), " "), "  ", " ")
}

func (q *QueryBuilder) Debug() string {
	return Debug(q.Stmt, q.args...)
}

func Debug(query string, args ...interface{}) string {
	s := CleanSQL(query)
	for k, v := range args {
		s = strings.Replace(s, fmt.Sprintf("$%d", k+1), fmt.Sprint(v), -1)
	}
	return s
}
