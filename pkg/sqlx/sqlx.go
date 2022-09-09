package sqlx

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/ckam225/golang/sqlx/internal/utils"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang.org/x/exp/maps"
)

type SQLX struct {
	dsn   string
	query string
	DB    *sqlx.DB
}

func Postgres(host string, port int, dbname, username, password, sslmode string, timeout int) (*SQLX, error) {
	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s connect_timeout=%d",
		host, port, dbname, username, password, sslmode, timeout,
	)
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error openning database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}
	return &SQLX{
		dsn: dsn,
		DB:  db,
	}, nil
}

func (s *SQLX) RunMigrations(migrationDir string) error {
	m, err := migrate.New(
		migrationDir,
		s.dsn,
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		return err
	}
	return nil
}

func (s *SQLX) RollbackMigrations(migrationDir string) error {
	m, err := migrate.New(migrationDir, s.dsn)
	if err != nil {
		return err
	}
	if err := m.Down(); err != nil {
		return err
	}
	return nil
}

func (s *SQLX) GetDNS() string {
	return s.dsn
}

func (s *SQLX) Close() error {
	return s.DB.Close()
}

func (s *SQLX) Select(table string, fields ...string) *SQLX {
	var f string
	if len(fields) == 0 {
		f = "*"
	} else {
		f = strings.Join(fields, ",")
	}
	s.query = fmt.Sprintf(`SELECT %s FROM %s`, f, table)
	return s
}

func (s *SQLX) Where(field, condition string, value interface{}) *SQLX {
	s.query = normalizeWhereQuery(s.query, "AND", field, condition, value)
	return s
}

func (s *SQLX) WhereIn(field string, values interface{}) *SQLX {
	s.query = normalizeWhereInQuery(s.query, "AND", field, "IN", values)
	return s
}

func (s *SQLX) WhereNotIn(field string, values interface{}) *SQLX {
	s.query = normalizeWhereInQuery(s.query, "AND", field, "NOT IN", values)
	return s
}

func (s *SQLX) OrWhere(field, condition string, value interface{}) *SQLX {
	s.query = normalizeWhereQuery(s.query, "OR", field, condition, value)
	return s
}

func (s *SQLX) OrWhereIn(field string, values interface{}) *SQLX {
	s.query = normalizeWhereInQuery(s.query, "OR", field, "IN", values)
	return s
}

func (s *SQLX) OrWhereNotIn(field string, values interface{}) *SQLX {
	s.query = normalizeWhereInQuery(s.query, "OR", field, "NOT IN", values)
	return s
}

func (s *SQLX) Offset(offset uint) *SQLX {
	s.query += fmt.Sprintf(" OFFSET %d", offset)
	return s
}

func (s *SQLX) Limit(limit uint) *SQLX {
	s.query += fmt.Sprintf(" LIMIT %d", limit)
	return s
}

func (s *SQLX) OrderBy(fields ...string) *SQLX {
	s.query += " ORDER BY " + strings.Join(fields, ",")
	return s
}

func (s *SQLX) Desc(fields ...string) *SQLX {
	s.query += " DESC"
	return s
}

func (s *SQLX) Asc(fields ...string) *SQLX {
	s.query += " ASC"
	return s
}

func (s *SQLX) GroupBy(fields ...string) *SQLX {
	s.query += " GROUP BY " + strings.Join(fields, ",")
	return s
}

func (s *SQLX) Values(table string, fields ...string) *SQLX {
	s.query = "INSERT INTO"
	if len(fields) > 0 {
		s.query += fmt.Sprintf("(%s)", strings.Join(fields, ","))
	}
	return s
}

func (s *SQLX) Insert(args ...interface{}) error {
	s.query = normalizeInsertValueQuery(s.query, args)
	_, err := s.DB.Exec(s.query, args...)
	return err
}

func (s *SQLX) Create(table string, obj interface{}) string {
	attrs := maps.Keys(utils.StructToMap(&obj))
	query := fmt.Sprintf(`INSERT INTO %s(%s) VALUES (:%s)`,
		table,
		strings.Join(attrs, ","),
		strings.Join(attrs, ",:"),
	)
	//_, err := s.DB.NamedExec(query, obj)
	return query
}

func normalizeInsertValueQuery(query string, fields ...interface{}) string {
	q := []string{}
	for k := range fields {
		q = append(q, fmt.Sprint("$", k+1))
	}
	return query + fmt.Sprintf(" VALUES(%s)", strings.Join(q, ","))
}

func (s *SQLX) Get(dest interface{}) error {
	stmt, err := s.DB.Preparex(s.query)
	if err != nil {
		return err
	}
	if err := stmt.Select(dest); err != nil {
		return err
	}
	return nil
}

func (s *SQLX) RawSQL(dest interface{}, query string, args ...string) *SQLX {
	s.query = query
	return s
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
