package main

import (
	"fmt"

	"github.com/ckam225/golang/sqlx/database"
	"github.com/ckam225/golang/sqlx/entity"
	"github.com/ckam225/golang/sqlx/repository"
)

func main() {
	repo, err := repository.New(&database.Config{
		Host:     "host.docker.internal",
		Port:     "5432",
		Username: "postgres",
		Password: "postgres",
		Database: "golang_sqlx",
		SSLmode:  "disable",
		Timeout:  3000,
	})
	if err != nil {
		fmt.Printf(err.Error())
		panic(err)
	}
	CreateOne(repo)

	InsertOne(repo)

}

func FetchList(repo *repository.Repository) {
	people, _ := repo.GetPersons(100, 0)
	jason, john := people[0], people[1]

	fmt.Printf("%#v\n%#v", jason, john)
}

func FetchOne(repo *repository.Repository) {
	jason, err := repo.FindPerson(10)
	if err != nil {
		fmt.Printf(err.Error())
		panic(err)
	}
	fmt.Printf("%#v\n", jason)
}

func CreateOne(repo *repository.Repository) {
	//p, err := repo.CreatePerson(entity.Person{FirstName: "Wour", LastName: "Carlos", Email: "aecarlos@ab.co.nz"})
	p, err := repo.CreatePersonWithPrepare(entity.Person{FirstName: "Wour", LastName: "Carlos", Email: "aecarlos@ab.co.nz"})
	if err != nil {
		fmt.Printf(err.Error())
		panic(err)
	}
	fmt.Printf("%#v\n", p)
	fmt.Println(p.ID, p.FirstName, p.Email)
}

func InsertOne(repo *repository.Repository) {
	//err := repo.InsertNamedPerson(entity.Person{FirstName: "Ngani", LastName: "Laumape", Email: "nlaumape@ab.co.nz"})
	err := repo.InsertPersonWithPrepare(entity.Person{FirstName: "Ngani", LastName: "Laumape", Email: "nlaumape@ab.co.nz"})
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func BatchInsert(repo *repository.Repository) {
	personStructs := []entity.Person{
		{FirstName: "Ardie", LastName: "Savea", Email: "asavea@ab.co.nz"},
		{FirstName: "Sonny Bill", LastName: "Williams", Email: "sbw@ab.co.nz"},
		{FirstName: "Ngani", LastName: "Laumape", Email: "nlaumape@ab.co.nz"},
	}
	err := repo.BatchInsertPerson(personStructs)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
