package sqlx

import (
	"fmt"
	"strings"
	"time"

	"github.com/ckam225/golang/sqlx/internal/utils"
	"golang.org/x/exp/maps"
)

func (s *Database) Create(table string, obj interface{}) error {
	data := utils.ConvertStructToMap(obj)
	delete(data, "id")
	createdAt := data["created_at"]
	if createdAt != nil {
		data["created_at"] = time.Now()
	}
	updatedAt := data["updated_at"]
	if updatedAt != nil {
		data["updated_at"] = time.Now()
	}
	keys := maps.Keys(data)
	var values []interface{}
	var query strings.Builder
	query.WriteString(fmt.Sprintf(`insert into %s(%s)`,
		table,
		strings.Join(keys, ","),
	))
	query.WriteString(" values(")

	for k, v := range keys {
		query.WriteString(fmt.Sprintf("$%d", k+1))
		values = append(values, data[v])
		if k < len(keys)-1 {
			query.WriteString(",")
		}
	}
	query.WriteString(") returning *")
	return s.DB.Get(obj, query.String(), values...)
}

func (s *Database) Save(table string, obj interface{}) error {
	data := utils.ConvertStructToMap(&obj)
	ID, ok := data["id"]
	if !ok {
		return fmt.Errorf("model must contain id")
	}
	delete(data, "id")

	createdAt := data["created_at"]
	if createdAt != nil {
		delete(data, "created_at")
	}
	updatedAt := data["updated_at"]
	if updatedAt != nil {
		data["updated_at"] = time.Now()
	}
	keys := maps.Keys(data)
	var args []interface{}

	var query strings.Builder
	query.WriteString("update " + table + " set ")
	for i, k := range keys {
		query.WriteString(fmt.Sprintf("%s=$%d", k, i+1))
		if i < len(keys)-1 {
			query.WriteString(",")
		}
		args = append(args, data[k])
	}
	args = append(args, ID)
	query.WriteString(fmt.Sprintf(" where id=$%d  returning *", len(args)))
	for i, k := range keys {
		fmt.Printf("%d %s, %v\n", i, k, args[i])
	}
	return s.DB.Get(obj, query.String(), args...)
}
