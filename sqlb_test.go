package sqlb

import (
	"testing"
)

func TestSqlBuilder(t *testing.T) {

	builder := QueryBuilder{
		Stmt:    "select * from users",
		OrderBy: "ID",
		Limit:   90,
		Offset:  7,
	}
	builder.Where("id", "=", 1).
		Or("email", "=", "aaa@ajks.com").
		Where("age", "in", 30, 67, "80080").
		Or("role", "in", "admin", "driver").
		GroupBy("id", "age").
		Having("email", "=", "aaa@ajks.com").
		Or("item", "in", 0, 1).
		Build()
	result := `select * from users where id = 1 or email = aaa@ajks.com and age in (30,67,80080) or role in (admin,driver) group by id,age having email = aaa@ajks.com or item in (0,10) order by ID limit 90 offset 7;`
	if result != builder.Debug() {
		t.Errorf("Expected: %s , got: %s", result, builder.Debug())
	}
}
