package users

import (
	"e-library/books"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBooks(t *testing.T) {

	//testing canBorrowMore function
	t.Run("Checking borrow more books than 3", func(t *testing.T) {
		res := canBorrowMore(3)
		assert.Equal(t, true, res, "Should be true")
	})
	t.Run("Checking borrow more books than 3", func(t *testing.T) {
		res := canBorrowMore(4)
		assert.Equal(t, false, res, "Should be false")
	})

	//testing checkAlreadyBorrowed function
	t.Run("Checking already borrowed book which intending borrow again", func(t *testing.T) {
		err := checkAlreadyBorrowed(3, 1)
		assert.Equal(t, ErrorAlreadyBorrowedBook, err, "Should be error")
	})
	t.Run("Checking borrow book which not already borrowed", func(t *testing.T) {
		err := checkAlreadyBorrowed(1, 1)
		assert.Nil(t, err)
	})

	//testing borrowBook function
	t.Run("Checking borrow book which allowed", func(t *testing.T) {
		err := borrowBook(5, 20)
		assert.Nil(t, err)
	})
	t.Run("Checking borrow book which borrowed already from anyone", func(t *testing.T) {
		err := borrowBook(5, 6)
		assert.Equal(t, books.ErrorUnavailableBook, err, "Should be error")
	})
	t.Run("Checking borrow same book which already borrowed", func(t *testing.T) {
		err := borrowBook(5, 20)
		assert.Equal(t, books.ErrorUnavailableBook, err, "Should be error")
	})
	t.Run("Checking borrow more than 3", func(t *testing.T) {
		err := borrowBook(4, 19)
		assert.Equal(t, ErrorBorrowLimit, err, "Should be error")
	})

	//testing return book
	t.Run("Checking return book which allowed", func(t *testing.T) {
		err := returnBook(5, 20)
		assert.Nil(t, err)
	})
	t.Run("Checking return book which not borrowed", func(t *testing.T) {
		err := returnBook(5, 5)
		assert.Equal(t, ErrorReturnWrongBook, err, "Should be error")
	})
}
