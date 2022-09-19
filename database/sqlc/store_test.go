package database

import (
	"context"
	"database/sql"
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestStoreTransaction(t *testing.T) {
	ctx := context.Background()
	err := store.Transaction(ctx, func(q *Queries) error {
		author, err := q.CreateAuthor(ctx, CreateAuthorParams{
			Name: faker.Name(),
			Bio:  faker.Word() + " " + faker.Word(),
		})
		assert.NotEmpty(t, author)
		if err != nil {
			return err
		}
		book, err := q.CreateBook(ctx, CreateBookParams{
			Title:    faker.DomainName(),
			AuthorID: sql.NullInt32{Int32: author.ID, Valid: true},
		})
		if err != nil {
			return err
		}
		assert.NotEmpty(t, book)
		return nil
	})
	assert.NoError(t, err)
}
