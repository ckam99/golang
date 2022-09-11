package sqlx

import (
	"database/sql"
	"fmt"
	"strings"
)

type InsertQuery struct {
	Base   *Database
	Table  string
	Fields []string
	Values []interface{}
}

func (db *Database) Insert(t string, fields ...string) *InsertQuery {
	return &InsertQuery{
		Table:  t,
		Fields: fields,
		Base:   db,
	}
}

func (i *InsertQuery) Exec(args ...interface{}) (sql.Result, error) {
	var q strings.Builder
	q.WriteString("insert into " + i.Table)
	if len(i.Fields) > 0 {
		q.WriteString(fmt.Sprintf("(%s)", strings.Join(i.Fields, ",")))
	}
	q.WriteString(" values(")
	for v := range args {
		q.WriteString(fmt.Sprintf("$%d", v+1))
		if v < len(args)-1 {
			q.WriteString(",")
		}
	}
	q.WriteString(")")
	fmt.Println(q.String())
	return i.Base.DB.Exec(q.String(), args...)
}
