package main

import (
	"fmt"

	"github.com/ckam225/golang/sqlx/pkg/sqlx"
)

func main() {
	dns, err := sqlx.Postgres(
		"host.docker.internal",
		5432,
		"golang_sqlx",
		"postgres",
		"postgres",
		"disable",
		3000,
	)
	defer dns.Close()

	if err != nil {
		fmt.Printf(err.Error())
		panic(err)
	}
	// p := []entity.Person{}
	// p := entity.Person{}

	// p := entity.Person{
	// 	// ID:        7,
	// 	FirstName: faker.FirstName(),
	// 	LastName:  faker.LastName(),
	// 	Email:     faker.Email(),
	// }

	// p := entity.Person{
	// 	ID:        12,
	// 	FirstName: "Doe",
	// 	LastName:  "John",
	// 	Email:     "john.doe@mail.ru",
	// }

	result, err := dns.Update("persons", "firstname", "lastname", "email").
		Where("id", "=", 2).
		Exec("toto", "tata", "tota@example.com")
	if err != nil {
		panic(err)
	}

	fmt.Printf("after: %v\n", result)

	// fmt.Println(dns.Select("persons").Where("id", "=", 2).Get(&p))

}
