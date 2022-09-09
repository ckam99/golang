package repository

import (
	"fmt"

	"github.com/ckam225/golang/sqlx/internal/entity"
	"github.com/jmoiron/sqlx"
)

const (
	personTable = "persons"
)

type PersonRepository struct {
	*sqlx.DB
}

type IPersonRepository interface {
	GetPersons(limit, offset int) ([]entity.Person, error)
	FindPerson(id int) (*entity.Person, error)
	CreatePerson(p entity.Person) (*entity.Person, error)
	InsertPerson(p entity.Person) (int, error)
	InsertNamedPerson(person entity.Person) error
	BatchInsertPerson(persons []entity.Person) error
	CreatePersonWithPrepare(person entity.Person) (*entity.Person, error)
	InsertPersonWithPrepare(person entity.Person) error

	Count(id int) (int, error)
}

func NewPersonRepository(db *sqlx.DB) IPersonRepository {
	return &PersonRepository{
		DB: db,
	}
}

func (r *PersonRepository) GetPersons(limit, offset int) ([]entity.Person, error) {
	query := fmt.Sprintf(`SELECT * FROM %s LIMIT $1 OFFSET $2`, personTable)
	people := []entity.Person{}
	if err := r.Select(&people, query, limit, offset); err != nil {
		return []entity.Person{}, err
	}
	return people, nil
}

func (r *PersonRepository) FindPerson(id int) (*entity.Person, error) {
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, personTable)
	person := entity.Person{}
	if err := r.Get(&person, query, id); err != nil {
		return nil, fmt.Errorf("[error finding user] %w\n", err)
	}
	return &person, nil
}

func (s *PersonRepository) CreatePerson(t entity.Person) (*entity.Person, error) {
	query := fmt.Sprintf(`INSERT INTO %s (first_name, last_name, email) VALUES ($1, $2, $3) RETURNING *`, personTable)
	if err := s.Get(&t, query, t.FirstName, t.LastName, t.Email); err != nil {
		return nil, fmt.Errorf("[error creating user] %w", err)
	}
	return &t, nil
}

func (r *PersonRepository) InsertPerson(person entity.Person) (int, error) {
	query := fmt.Sprintf(`INSERT INTO %s (first_name, last_name, email) VALUES ($1, $2, $3)`, personTable)
	rows, err := r.MustExec(query, person.FirstName, person.LastName, person.Email).RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("[error inserting user] %w", err)
	}
	return int(rows), nil
}

func (r *PersonRepository) InsertNamedPerson(person entity.Person) error {
	query := fmt.Sprintf(`INSERT INTO %s (first_name, last_name, email) VALUES (:first_name, :last_name, :email)`, personTable)
	// Named queries can use structs, so if you have an existing struct (i.e. person := &Person{}) that you have populated, you can pass it in as &person
	_, err := r.NamedExec(query, &person)
	return err
}

func (r *PersonRepository) CreatePersonWithPrepare(person entity.Person) (*entity.Person, error) {
	stmt, err := r.Preparex(fmt.Sprintf(`INSERT INTO %s (first_name, last_name, email) VALUES ($1, $2, $3) RETURNING *`, personTable))
	if err != nil {
		return nil, err
	}
	if err := stmt.Get(&person, person.FirstName, person.LastName, person.Email); err != nil {
		return nil, err
	}

	// stmt, err := r.Prepare(fmt.Sprintf(`INSERT INTO %s (first_name, last_name, email) VALUES ($1, $2, $3) RETURNING *`, personTable))
	// if err != nil {
	// 	return nil, err
	// }
	// row := stmt.QueryRow(person.FirstName, person.LastName, person.Email)
	// if err := row.Scan(&person); err != nil {
	// 	return nil, err
	// }

	return &person, nil
}

func (r *PersonRepository) InsertPersonWithPrepare(person entity.Person) error {
	stmt, err := r.Prepare(fmt.Sprintf(`INSERT INTO %s (first_name, last_name, email) VALUES ($1, $2, $3)`, personTable))
	if err != nil {
		return nil
	}
	_, err = stmt.Exec(person.FirstName, person.LastName, person.Email)
	return err
}

func (r *PersonRepository) BatchInsertPerson(persons []entity.Person) error {
	query := fmt.Sprintf(`INSERT INTO %s (first_name, last_name, email) VALUES (:first_name, :last_name, :email)`, personTable)
	_, err := r.NamedExec(query, persons)
	return err
}

func (r *PersonRepository) BulkInsertPerson(persons []entity.Person) error {
	query := fmt.Sprintf(`INSERT INTO %s (first_name, last_name, email) VALUES (:first_name, :last_name, :email)`, personTable)
	tx := r.MustBegin()
	for person := range persons {
		tx.NamedExec(query, &person)
	}
	return tx.Commit()
}

func (r *PersonRepository) Count(id int) (int, error) {
	var count int
	query := fmt.Sprintf(`SELECT COUNT(*) FROM %s WHERE id = $1`, personTable)
	if err := r.Get(&count, query, id); err != nil {
		return count, err
	}
	return count, nil
}
