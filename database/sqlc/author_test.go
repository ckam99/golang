package database

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/assert"
)

func createRandomAuthor(t *testing.T) Author {
	payload := CreateAuthorParams{
		Name: faker.Name(),
		Bio:  faker.Paragraph(),
	}
	author, err := store.CreateAuthor(context.Background(), payload)
	assert.NoError(t, err)
	assert.NotEmpty(t, author)

	assert.Equal(t, payload.Name, author.Name)
	assert.Equal(t, payload.Bio, author.Bio)
	assert.NotZero(t, author.ID)
	assert.NotZero(t, author.CreatedAt)
	assert.Zero(t, author.UpdatedAt)

	return author
}

func TestCreateAuthor(t *testing.T) {
	createRandomAuthor(t)
}

func TestGetAuthor(t *testing.T) {
	author1 := createRandomAuthor(t)
	author, err := store.GetAuthor(context.Background(), author1.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, author)
	assert.Equal(t, author1.Name, author.Name)
	assert.Equal(t, author1.Bio, author.Bio)
	assert.Equal(t, author1.ID, author.ID)
	assert.WithinDuration(t, author1.CreatedAt.Time, author.CreatedAt.Time, time.Second)

}

func TestUpdateAuthor(t *testing.T) {
	oldAuthor := createRandomAuthor(t)
	author, err := store.UpdateAuthor(context.Background(), UpdateAuthorParams{
		ID:   oldAuthor.ID,
		Name: sql.NullString{Valid: true, String: faker.Name()},
		Bio:  sql.NullString{Valid: true, String: oldAuthor.Bio},
	})
	assert.NoError(t, err)
	assert.NotEqual(t, oldAuthor.Name, author.Name)
	assert.Equal(t, oldAuthor.Bio, author.Bio)
	assert.Equal(t, oldAuthor.ID, author.ID)
	assert.WithinDuration(t, oldAuthor.CreatedAt.Time, author.CreatedAt.Time, time.Second)

	assert.NotZero(t, author.UpdatedAt)
}

func TestUpdatePartialAuthor(t *testing.T) {
	oldAuthor := createRandomAuthor(t)
	newName := faker.Name()
	author, err := store.UpdateAuthor(context.Background(), UpdateAuthorParams{
		ID:   oldAuthor.ID,
		Name: sql.NullString{Valid: true, String: newName},
	})
	assert.NoError(t, err)
	assert.NotEqual(t, oldAuthor.Name, author.Name)
	assert.Equal(t, newName, author.Name)
	assert.Equal(t, oldAuthor.Bio, author.Bio)
	assert.WithinDuration(t, oldAuthor.CreatedAt.Time, author.CreatedAt.Time, time.Second)
	assert.NotZero(t, author.UpdatedAt)
}

func TestDeleteAuthor(t *testing.T) {
	author := createRandomAuthor(t)
	err := store.DeleteAuthor(context.Background(), author.ID)
	assert.NoError(t, err)

	obj, err := store.GetAuthor(context.Background(), author.ID)
	assert.Error(t, err)
	assert.EqualError(t, err, sql.ErrNoRows.Error())
	assert.Empty(t, obj)
}

func TestGetAllAuthors(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAuthor(t)
	}
	args := GetAllAuthorsParams{
		Limit:  5,
		Offset: 4,
	}
	authors, err := store.GetAllAuthors(context.Background(), args)
	assert.NoError(t, err)
	assert.Len(t, authors, 5)
	for _, v := range authors {
		assert.NotEmpty(t, v)
	}
}
