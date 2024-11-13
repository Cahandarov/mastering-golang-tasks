package unit_testing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculator(t *testing.T) {
	t.Run("Testing addition", func(t *testing.T) {
		result, err := Calculate(2, 5, "+")
		assert.Nil(t, err)
		assert.Equal(t, 7, result, "they should be equal")
	})

	t.Run("Testing subtraction", func(t *testing.T) {
		result, err := Calculate(5, 5, "-")
		assert.Nil(t, err)
		assert.Equal(t, 0, result, "they should be equal")
	})

	t.Run("Testing multiplication", func(t *testing.T) {
		result, err := Calculate(2, 5, "*")
		assert.Nil(t, err)
		assert.Equal(t, 10, result, "they should be equal")
	})

	t.Run("Testing division", func(t *testing.T) {
		result, err := Calculate(10, 5, "/")
		assert.Nil(t, err)
		assert.Equal(t, 2, result, "they should be equal")
	})

	t.Run("Testing division by zero", func(t *testing.T) {
		_, err := Calculate(10, 0, "/")
		assert.Equal(t, ErrorZeroDivide, err, "Should return error: divide by zero")
	})

	t.Run("Testing unknown operation", func(t *testing.T) {
		_, err := Calculate(10, 2, "8")
		assert.Equal(t, ErrorUnknownInput, err, "Should return error: unknown input")
	})
}
