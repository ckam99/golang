package author

import (
	"app/pkg/clients/postgresql"
	"context"
	"log"
	"os"
	"testing"

	"github.com/bxcodec/faker/v4"
)

var serv Service

func TestMain(t *testing.M) {
	pg, err := postgresql.NewClient(context.TODO(), postgresql.Config{
		Host:     "host.docker.internal",
		Port:     "54323",
		Database: "demo",
		Password: "postgres",
		Username: "postgres",
	}, 3)
	if err != nil {
		log.Fatal("connot connect to the database: ", err.Error())
	}
	serv = NewService(pg)
	os.Exit(t.Run())

	// tearclean
	log.Println("Clean database")
	if _, err = pg.Exec(context.TODO(), "truncate table authors cascade;"); err != nil {
		log.Fatal("Error: ", postgresql.Error(err))
	}

}

func CreateRandomAuhtor(t *testing.T) Author {
	author := Author{
		Name:      faker.Name(),
		Biography: faker.Word(),
	}
	if err := serv.CreateAuthor(context.TODO(), &author); err != nil {
		t.Errorf("createAuhtor: expected: nil got: %s", err.Error())
	}
	if author.ID == 0 {
		t.Errorf("createAuhtor: author ID,  expected: > 0  got: 0")
	}
	// if !author.UpdatedAt.Time.IsZero() {
	// 	t.Errorf("createAuhtor[updated_at should be zero] expected : true, got: false")
	// }
	// if author.CreatedAt.Time.IsZero() {
	// 	t.Errorf("createAuhtor[created_at not zero] expected : true, got: false")
	// }
	return author
}

func TestCreateAuhtorService(t *testing.T) {
	CreateRandomAuhtor(t)
}

func TestGetAuhtorService(t *testing.T) {
	author1 := CreateRandomAuhtor(t)
	author2, err := serv.FindAuthor(context.TODO(), author1.ID)
	if err != nil {
		t.Errorf("TestGetAuhtor: expected: nil got: %s", err.Error())
	}
	if author1.ID != author2.ID || author1.Name != author2.Name || author1.Biography != author2.Biography {
		t.Errorf("TestGetAuhtor: author1 should be equal to author2 expected: true, got: false")
	}
}

func TestGetAllAuthorsService(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomAuhtor(t)
	}
	authors, err := serv.GetAuthors(context.TODO(), FilterParamsDTO{
		Limit:  5,
		Offset: 5,
	})
	if err != nil {
		t.Errorf("TestGetAllAuhtors: expected: nil got: %s", err.Error())
	}
	if len(authors) != 5 {
		t.Errorf("TestGetAllAuhtors(authors size) expected 5, got: %d", len(authors))
	}
}

func TestUpdateAuhtorService(t *testing.T) {
	author := CreateRandomAuhtor(t)
	author1 := Author{
		ID:        author.ID,
		Name:      faker.Name(),
		Biography: faker.Sentence(),
	}
	if err := serv.UpdateAuthor(context.TODO(), &author1); err != nil {
		t.Errorf("TestGetAuhtor: expected: nil got: %s", err.Error())
	}
	if !author1.UpdatedAt.Valid {
		t.Errorf("TestGetAuhtor[updated_at not zero] expected : true, got: false")
	}
}

func TestDeleteAuhtorService(t *testing.T) {
	author := CreateRandomAuhtor(t)
	if err := serv.DeleteAuthor(context.TODO(), author.ID, false); err != nil {
		t.Errorf("TestDeleteAuhtor: expected: nil got: %s", err.Error())
	}
	author2, err := serv.FindAuthor(context.TODO(), author.ID)
	if err == nil {
		t.Errorf("TestGetAuhtor: expected: nil got: %s", err.Error())
	}
	if author.ID == author2.ID || author.Name == author2.Name {
		t.Errorf("TestGetAuhtor: author1 should be equal to author2 expected: true, got: false")
	}
}

// func TestSoftDeleteAuhtorService(t *testing.T) {
// 	author := CreateRandomAuhtor(t)
// 	if err := serv.SoftDeleteAuthor(context.TODO(), &author); err != nil {
// 		t.Errorf("TestDeleteAuhtor: expected: nil got: %s", err.Error())
// 	}
// 	if !author.DeletedAt.Valid {
// 		t.Errorf("TestDeleteAuhtor[deleted_at not zero] expected : true, got: false")
// 	}
// }
