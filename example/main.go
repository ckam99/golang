package main

import (
	"fmt"

	"github.com/ckam225/sqlb"
)

func main() {
	builder := sqlb.QueryBuilder{
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
	fmt.Println(builder.Debug())
}
