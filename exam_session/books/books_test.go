package books

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBook(t *testing.T) {
	t.Run("Adding new book", func(t *testing.T) {
		err := AddBook(21, "New book for testing", "Shamkhal")
		assert.Equal(t, nil, err, "Should be add new book")
		assert.Nil(t, err)
	})
	t.Run("Adding new book wrong id", func(t *testing.T) {
		err := AddBook(-21, "New book for testing", "Shamkhal")
		assert.Equal(t, ErrorOnAddingNewBook, err, "Should be error")
	})
	t.Run("Adding new book with empty title ", func(t *testing.T) {
		err := AddBook(22, "", "Shamkhal")
		assert.Equal(t, ErrorOnAddingNewBook, err, "Should be error")
	})
	t.Run("Adding new book with empty author", func(t *testing.T) {
		err := AddBook(23, "New book for testing", "")
		assert.Equal(t, ErrorOnAddingNewBook, err, "Should be error")
	})

	t.Run("Check available book with available id", func(t *testing.T) {
		res, err := BookAvailable(3)
		assert.Equal(t, true, res, "Should be return true")
		assert.Nil(t, err)
	})
	t.Run("Check available book with borrowed book id", func(t *testing.T) {
		res, err := BookAvailable(1)
		assert.Equal(t, false, res, "Should be return false")
		assert.Equal(t, ErrorUnavailableBook, err, "Should be return false")
	})
	t.Run("Check available book with wrong id", func(t *testing.T) {
		res, err := BookAvailable(100)
		assert.Equal(t, false, res, "Should be return false")
		assert.Nil(t, err)
	})
}
