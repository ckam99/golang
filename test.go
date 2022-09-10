package main

import "github.com/ckam225/golang/sqlx/pkg/orm"

// func main() {
// 	// dns, err := sqlx.Postgres(
// 	// 	"host.docker.internal",
// 	// 	5432,
// 	// 	"golang_sqlx",
// 	// 	"postgres",
// 	// 	"postgres",
// 	// 	"disable",
// 	// 	3000,
// 	// )
// 	// defer dns.Close()

// 	// if err != nil {
// 	// 	fmt.Printf(err.Error())
// 	// 	panic(err)
// 	// }

// 	// // if err := dns.Select("persons").Get(&persons); err != nil {
// 	// // 	panic(err)
// 	// // }

// 	// // fmt.Println(persons)

// 	// pers := entity.Person{
// 	// 	// FirstName: faker.FirstName(),
// 	// 	// LastName:  faker.LastName(),
// 	// 	// Email:     faker.Email(),
// 	// }

// 	// dns.Create("persons", pers)

// 	// p := []entity.Person{}
// 	// fmt.Println(dns.Select("persons").Where("id", "=", 2).Get(&p))

// }

func main() {
	db := orm.Database{}

	//db.Insert("users", "id", "age", "name").Exec(34, 565, "name")

	//db.Update("users", "id", "age", "name").Exec(34, 565, "name")

	db.Select("users").
		Where("name", "=", "Angel").
		OrWhere("email", "like", "'%@yandex.ru'").
		Where("id", ">", 1).
		WhereIn("age", []int{4, 10, 99}).
		OrWhereIn("lastname", []string{"jean", "rachel"}).
		GroupBy("id").
		OrderBy("email").
		Desc().
		Limit(1000).
		Offset(10).
		Get()

}
