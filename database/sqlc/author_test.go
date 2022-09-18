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
		Fullname: faker.Name(),
		Bio:      faker.Paragraph(),
	}
	author, err := testQueries.CreateAuthor(context.Background(), payload)
	assert.NoError(t, err)
	assert.NotEmpty(t, author)

	assert.Equal(t, payload.Fullname, author.Fullname)
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
	author, err := testQueries.GetAuthor(context.Background(), author1.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, author)
	assert.Equal(t, author1.Fullname, author.Fullname)
	assert.Equal(t, author1.Bio, author.Bio)
	assert.Equal(t, author1.ID, author.ID)
	assert.WithinDuration(t, author1.CreatedAt.Time, author.CreatedAt.Time, time.Second)

}

func TestUpdateAuthor(t *testing.T) {
	author1 := createRandomAuthor(t)
	payload := UpdateAuthorParams{
		ID:       author1.ID,
		Fullname: faker.Name(),
		Bio:      author1.Bio,
	}
	author, err := testQueries.UpdateAuthor(context.Background(), payload)
	assert.NoError(t, err)
	assert.NotEmpty(t, author)
	assert.NotEqual(t, author1.Fullname, author.Fullname)
	assert.Equal(t, author1.Bio, author.Bio)
	assert.Equal(t, author1.ID, author.ID)
	assert.WithinDuration(t, author1.CreatedAt.Time, author.CreatedAt.Time, time.Second)

	assert.Equal(t, payload.Fullname, author.Fullname)
	assert.NotZero(t, author.UpdatedAt)
}

func TestDeleteAuthor(t *testing.T) {
	author := createRandomAuthor(t)
	err := testQueries.DeleteAuthor(context.Background(), author.ID)
	assert.NoError(t, err)

	obj, err := testQueries.GetAuthor(context.Background(), author.ID)
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
	authors, err := testQueries.GetAllAuthors(context.Background(), args)
	assert.NoError(t, err)
	assert.Len(t, authors, 5)
	for _, v := range authors {
		assert.NotEmpty(t, v)
	}

}
