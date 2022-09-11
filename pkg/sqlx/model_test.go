package sqlx

import (
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/ckam225/golang/sqlx/internal/entity"
	"github.com/stretchr/testify/assert"
)

var db *Database

func ConntectDb() {
	dns, err := Postgres(
		"host.docker.internal",
		5432,
		"golang_sqlx",
		"postgres",
		"postgres",
		"disable",
		3000,
	)
	if err != nil {
		panic(err.Error())
	}
	defer dns.Close()

	db = dns
}

func TestModelCreate(t *testing.T) {
	person := entity.Person{
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Email:     faker.Email(),
	}
	err := db.Create("persons", &person)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, person.ID)
}

func TestModelSave(t *testing.T) {
	person := entity.Person{
		ID:        1,
		FirstName: "Doe",
		LastName:  "John",
		Email:     "john.doe@mail.ru",
	}
	err := db.Save("persons", &person)
	assert.Nil(t, err)
	assert.NotEqual(t, "John", person.LastName)
}
