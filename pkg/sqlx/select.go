package sqlx

import (
	"fmt"
	"strings"
)

type SelectQuery struct {
	table          string
	columns        []string
	whereQueries   []whereQuery
	groupByColumns []string
	orderColumns   []string
	ascending      string
	Take           uint
	Skip           uint
}

// func (s *SelectQuery) addWhereQuery(column, oper string, value interface{}, clause string) {
// 	s.whereQueries = append(s.whereQueries, whereQuery{
// 		Column:   column,
// 		Value:    value,
// 		Operator: oper,
// 		Type:     clause,
// 	})
// }

func (db *Database) Select(table string, columns ...string) *SelectQuery {
	return &SelectQuery{
		table:   table,
		columns: columns,
	}
}

func (q *SelectQuery) GroupBy(columns ...string) *SelectQuery {
	q.groupByColumns = columns
	return q
}

func (q *SelectQuery) OrderBy(columns ...string) *SelectQuery {
	q.orderColumns = columns
	return q
}

func (q *SelectQuery) Desc() *SelectQuery {
	q.ascending = "desc"
	return q
}

func (q *SelectQuery) Asc() *SelectQuery {
	q.ascending = "asc"
	return q
}

func (q *SelectQuery) Limit(limit uint) *SelectQuery {
	q.Take = limit
	return q
}

func (q *SelectQuery) Offset(offset uint) *SelectQuery {
	q.Skip = offset
	return q
}

func (q *SelectQuery) Where(column string, oper string, value interface{}) *SelectQuery {
	addWhereQuery(&q.whereQueries, column, oper, value, "and")
	return q
}

func (q *SelectQuery) OrWhere(column string, oper string, value interface{}) *SelectQuery {
	addWhereQuery(&q.whereQueries, column, oper, value, "or")
	return q
}

func (q *SelectQuery) WhereIn(column string, value interface{}) *SelectQuery {
	addWhereQuery(&q.whereQueries, column, "in", value, "and")
	return q
}

func (q *SelectQuery) WhereNotIn(column string, value interface{}) *SelectQuery {
	addWhereQuery(&q.whereQueries, column, "not in", value, "and")
	return q
}

func (q *SelectQuery) OrWhereIn(column string, value interface{}) *SelectQuery {
	addWhereQuery(&q.whereQueries, column, "in", value, "or")
	return q
}

func (q *SelectQuery) OrWhereNotIn(column string, value interface{}) *SelectQuery {
	addWhereQuery(&q.whereQueries, column, "not in", value, "or")
	return q
}

func (q *SelectQuery) Get() {
	var query strings.Builder
	var args []interface{}
	// select from
	query.WriteString("select ")
	if len(q.columns) > 0 {
		query.WriteString(strings.Join(q.columns, ","))
	} else {
		query.WriteString("*")
	}
	query.WriteString(" from " + q.table)
	// where
	if len(q.whereQueries) > 0 {
		sb := buildFilterQuery(q.whereQueries, &args)
		query.WriteString(sb)
	}
	// group by
	if len(q.groupByColumns) > 0 {
		//query.WriteString(" group by " + strings.Join(q.groupByColumns, ","))
		query.WriteString(" group by ")
		for k, v := range q.groupByColumns {
			query.WriteString(fmt.Sprintf("$%d", len(args)+1))
			if k < len(q.groupByColumns)-1 {
				query.WriteString(",")
			}
			args = append(args, v)
		}
	}

	// order by
	if len(q.orderColumns) > 0 {
		//query.WriteString(" order by " + strings.Join(q.orderColumns, ","))
		query.WriteString(" order by ")
		for k, v := range q.orderColumns {
			query.WriteString(fmt.Sprintf("$%d", len(args)+1))
			if k < len(q.orderColumns)-1 {
				query.WriteString(",")
			}
			args = append(args, v)
		}
	}

	// ascending
	if q.ascending != "" {
		query.WriteString(fmt.Sprintf(" $%d", len(args)+1))
		args = append(args, q.ascending)
	}

	// limit
	if q.Take > 0 {
		query.WriteString(fmt.Sprintf(" limit $%d", len(args)+1))
		args = append(args, q.Take)
	}
	// offset
	if q.Skip != 0 {
		query.WriteString(fmt.Sprintf(" offset $%d", len(args)+1))
		args = append(args, q.Skip)
	}

	fmt.Println("---query---")
	fmt.Println(query.String())
	fmt.Println("---args---")
	for _, v := range args {
		fmt.Printf("%v, ", v)
	}
	fmt.Println()

}
