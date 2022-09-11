package sqlx

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type UpdateQuery struct {
	instance     *Database
	Table        string
	Fields       []string
	whereQueries []whereQuery
}

func (db *Database) Update(table string, fields ...string) *UpdateQuery {
	return &UpdateQuery{
		Table:    table,
		Fields:   fields,
		instance: db,
	}
}

func (q *UpdateQuery) Where(column string, oper string, value interface{}) *UpdateQuery {
	addWhereQuery(&q.whereQueries, column, oper, value, "and")
	return q
}

func (q *UpdateQuery) OrWhere(column string, oper string, value interface{}) *UpdateQuery {
	addWhereQuery(&q.whereQueries, column, oper, value, "or")
	return q
}

func (q *UpdateQuery) WhereIn(column string, value interface{}) *UpdateQuery {
	addWhereQuery(&q.whereQueries, column, "in", value, "and")
	return q
}

func (q *UpdateQuery) WhereNotIn(column string, value interface{}) *UpdateQuery {
	addWhereQuery(&q.whereQueries, column, "not in", value, "and")
	return q
}

func (q *UpdateQuery) OrWhereIn(column string, value interface{}) *UpdateQuery {
	addWhereQuery(&q.whereQueries, column, "in", value, "or")
	return q
}

func (q *UpdateQuery) OrWhereNotIn(column string, value interface{}) *UpdateQuery {
	addWhereQuery(&q.whereQueries, column, "not in", value, "or")
	return q
}

func (q *UpdateQuery) Exec(args ...interface{}) (sql.Result, error) {
	query, err := q.getBaseQuery(&args)
	if err != nil {
		return nil, err
	}
	return q.instance.DB.Exec(query.String(), args...)
}

func (q *UpdateQuery) Get(dest interface{}, args ...interface{}) error {
	query, err := q.getBaseQuery(&args)
	if err != nil {
		return err
	}
	query.WriteString(" returning *")
	fmt.Println(query.String())
	return q.instance.DB.Get(dest, query.String(), args...)
}

func (q *UpdateQuery) getBaseQuery(args *[]interface{}) (*strings.Builder, error) {
	var query strings.Builder
	query.WriteString("update " + q.Table + " set ")
	fieldCount := len(q.Fields)
	if fieldCount == 0 || len(*args) == 0 {
		return nil, errors.New("update:fields and values need")
	}
	if fieldCount != len(*args) {
		return nil, errors.New("update:fields and values len don't match")
	}
	for i := 0; i < fieldCount; i++ {
		query.WriteString(fmt.Sprintf("%s=$%d", q.Fields[i], i+1))
		if i < fieldCount-1 {
			query.WriteString(",")
		}
	}
	// where
	if len(q.whereQueries) > 0 {
		sb := buildFilterQuery(q.whereQueries, args)
		query.WriteString(sb)
	}

	return &query, nil
}
