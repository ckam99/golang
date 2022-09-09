package main

import (
	"fmt"

	"github.com/bxcodec/faker/v3"
	"github.com/ckam225/golang/sqlx/internal/entity"
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

	// persons := []entity.Person{}

	// if err := dns.Select("persons").Get(&persons); err != nil {
	// 	panic(err)
	// }

	// fmt.Println(persons)

	p := entity.Person{
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Email:     faker.Email(),
	}
	fmt.Println(dns.Create("person", p))

}
