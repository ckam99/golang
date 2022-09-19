package database

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/assert"
)

func createRandomBook(t *testing.T) Book {
	author := createRandomAuthor(t)
	payload := CreateBookParams{
		AuthorID: sql.NullInt32{Int32: author.ID, Valid: true},
		Title:    faker.Sentence(),
	}
	book, err := store.CreateBook(context.Background(), payload)
	assert.NoError(t, err)
	assert.NotEmpty(t, author)

	assert.Equal(t, payload.Title, book.Title)
	assert.Equal(t, payload.AuthorID, book.AuthorID)
	assert.Equal(t, author.ID, book.AuthorID.Int32)
	assert.NotZero(t, book.ID)
	assert.NotZero(t, author.CreatedAt)
	assert.Zero(t, author.UpdatedAt)

	return book
}

func TestCreateBook(t *testing.T) {
	createRandomBook(t)
}

func TestGetBook(t *testing.T) {
	book := createRandomBook(t)
	book2, err := store.GetBook(context.Background(), book.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, book2)
	assert.Equal(t, book2.Title, book.Title)
	assert.Equal(t, book2.ID, book.ID)
	assert.WithinDuration(t, book2.CreatedAt.Time, book.CreatedAt.Time, time.Second)
}

func TestUpdateBook(t *testing.T) {
	book := createRandomBook(t)
	payload := UpdateBookParams{
		ID:       book.ID,
		Title:    faker.Sentence(),
		AuthorID: book.AuthorID,
	}
	err := store.UpdateBook(context.Background(), payload)
	assert.NoError(t, err)
	book2, err := store.GetBook(context.Background(), book.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, book2)
	assert.NotEqual(t, book2.Title, book.Title)
	assert.Equal(t, payload.Title, book2.Title)
	assert.NotZero(t, book2.UpdatedAt)
}

func TestDeleteBook(t *testing.T) {
	author := createRandomBook(t)
	err := store.DeleteBook(context.Background(), author.ID)
	assert.NoError(t, err)

	obj, err := store.GetBook(context.Background(), author.ID)
	assert.Error(t, err)
	assert.EqualError(t, err, sql.ErrNoRows.Error())
	assert.Empty(t, obj)
}

func TestGetAllBooks(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomBook(t)
	}
	args := GetAllBooksParams{
		Limit:  5,
		Offset: 4,
	}
	authors, err := store.GetAllBooks(context.Background(), args)
	assert.NoError(t, err)
	assert.Len(t, authors, 5)
	for _, v := range authors {
		assert.NotEmpty(t, v)
	}

}
